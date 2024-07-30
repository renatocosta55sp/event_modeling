package http

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"github.org/napp/product-management/internal/app/command"
	"github.org/napp/product-management/internal/app/query"
	"github.org/napp/product-management/internal/domain/product/events"
	"github.org/napp/product-management/internal/infra/adapters/persistence"
	"github.org/napp/product-management/internal/infra/config"
	"github.org/napp/product-management/internal/infra/service"
	"github.org/napp/product-management/pkg/building_blocks/infra/bus"
)

type HttpServer struct {
	Db *pgx.Conn
}

func (h HttpServer) CreateProduct(ctx *gin.Context) {

	createProductRequest := CreateProductRequest{}
	if err := ctx.ShouldBindJSON(&createProductRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var wg sync.WaitGroup
	eventBus := bus.NewEventBus()

	wireApp := service.NewApplication(ctx,
		&wg,
		eventBus,
		persistence.NewProductRepository(h.Db, config.Config.GetString("DB_SCHEMA")),
	)

	currentDateTime := time.Now().Format("2006-01-02 15:04:05.000")

	wg.Add(1)
	fmt.Println(createProductRequest.PriceFrom)
	_, err := wireApp.Commands.CreateProduct.Handle(ctx, command.CreateProductCommand{
		Name:       "Product 0102023",
		Code:       createProductRequest.Code,
		Stock:      createProductRequest.Stock,
		TotalStock: createProductRequest.TotalStock,
		CutStock:   createProductRequest.CutStock,
		PriceFrom:  createProductRequest.PriceFrom,
		PriceTo:    createProductRequest.PriceTo,
		CreatedBy:  123,
		UpdatedBy:  123,
		CreatedAt:  currentDateTime,
		UpdatedAt:  currentDateTime,
	})

	if err != nil {
		logrus.WithError(err).Error("failed to validate product on command")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	wg.Wait()
	if !eventBus.Raised(events.ProductCreatedEvent) {
		err := eventBus.GetError()
		logrus.WithError(err).Error("failed to create product on event handler")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})

}

func (h HttpServer) UpdateProduct(ctx *gin.Context) {

	code, _ := strconv.Atoi(ctx.Param("code"))

	updateProductRequest := UpdateProductRequest{}
	if err := ctx.ShouldBindJSON(&updateProductRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentDateTime := time.Now().Format("2006-01-02 15:04:05.000")

	var wg sync.WaitGroup
	eventBus := bus.NewEventBus()

	wireApp := service.NewApplication(ctx,
		&wg,
		eventBus,
		persistence.NewProductRepository(h.Db, config.Config.GetString("DB_SCHEMA")),
	)

	wg.Add(1)

	_, err := wireApp.Commands.UpdateProduct.Handle(ctx, command.UpdateProductCommand{
		Code:       code,
		Name:       updateProductRequest.Name,
		Stock:      updateProductRequest.Stock,
		TotalStock: updateProductRequest.TotalStock,
		CutStock:   updateProductRequest.CutStock,
		PriceFrom:  updateProductRequest.PriceFrom,
		PriceTo:    updateProductRequest.PriceTo,
		UpdatedBy:  12345,
		UpdatedAt:  currentDateTime,
	})
	if err != nil {
		logrus.WithError(err).Error("failed to validate product on command")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	wg.Wait()
	if !eventBus.Raised(events.ProductUpdatedEvent) {
		err := eventBus.GetError()
		logrus.WithError(err).Error("failed to update product on event handler")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})

}

func (h HttpServer) DeleteProduct(ctx *gin.Context) {

	code, _ := strconv.Atoi(ctx.Param("code"))

	var wg sync.WaitGroup
	eventBus := bus.NewEventBus()

	wireApp := service.NewApplication(ctx,
		&wg,
		eventBus,
		persistence.NewProductRepository(h.Db, config.Config.GetString("DB_SCHEMA")),
	)

	wg.Add(1)

	_, err := wireApp.Commands.DeleteProduct.Handle(ctx, command.DeleteProductCommand{
		Code: code,
	})
	if err != nil {
		logrus.WithError(err).Error("failed to validate product on command")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	wg.Wait()
	if !eventBus.Raised(events.ProductDeletedEvent) {
		err := eventBus.GetError()
		logrus.WithError(err).Error("failed to delete product on event handler")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})

}

func (h HttpServer) GetProducts(ctx *gin.Context) {

	var wg sync.WaitGroup
	eventBus := bus.NewEventBus()

	wireApp := service.NewApplication(ctx,
		&wg,
		eventBus,
		persistence.NewProductRepository(h.Db, config.Config.GetString("DB_SCHEMA")),
	)

	results, err := wireApp.Queries.AvailableProducts.Handle(ctx, query.AvailableProducts{})
	if err != nil {
		logrus.WithError(err).Error("failed to validate products on query")
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": results})

}
