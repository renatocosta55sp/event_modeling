package http

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.org/eventmodeling/product-management/internal/app/command"
	"github.org/eventmodeling/product-management/internal/app/query"
	"github.org/eventmodeling/product-management/internal/domain/product/events"
	"github.org/eventmodeling/product-management/internal/infra/adapters/persistence"
	"github.org/eventmodeling/product-management/internal/infra/config"
	"github.org/eventmodeling/product-management/internal/infra/service"
	"github.org/eventmodeling/product-management/pkg/building_blocks/infra/bus"
)

type HttpServer struct {
	Db *pgxpool.Pool
}

func (h HttpServer) CreateProduct(ctx *gin.Context) {

	requestData, exists := ctx.Get(requestDataKey)
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Request data not found"})
		return
	}

	data := requestData.(CreateProductRequest)

	var wg sync.WaitGroup
	eventBus := bus.NewEventBus()

	wireApp := service.NewApplication(ctx,
		&wg,
		eventBus,
		persistence.NewProductRepository(h.Db, config.Config.GetString("DB_SCHEMA")),
	)

	currentDateTime := time.Now().Format("2006-01-02 15:04:05.000")

	wg.Add(1)

	_, err := wireApp.Commands.CreateProduct.Handle(ctx, command.CreateProductCommand{
		Name:       "Product 0102023",
		Code:       data.Code,
		Stock:      data.Stock,
		TotalStock: data.TotalStock,
		CutStock:   data.CutStock,
		PriceFrom:  data.PriceFrom,
		PriceTo:    data.PriceTo,
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

	requestData, exists := ctx.Get(requestDataKey)
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Request data not found"})
		return
	}

	data := requestData.(UpdateProductRequest)

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
		Id:         data.Id,
		Name:       data.Name,
		Stock:      data.Stock,
		TotalStock: data.TotalStock,
		CutStock:   data.CutStock,
		PriceFrom:  data.PriceFrom,
		PriceTo:    data.PriceTo,
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

	requestData, exists := ctx.Get(requestDataKey)
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Request data not found"})
		return
	}

	data := requestData.(DeleteProductRequest)

	var wg sync.WaitGroup
	eventBus := bus.NewEventBus()

	wireApp := service.NewApplication(ctx,
		&wg,
		eventBus,
		persistence.NewProductRepository(h.Db, config.Config.GetString("DB_SCHEMA")),
	)

	wg.Add(1)

	_, err := wireApp.Commands.DeleteProduct.Handle(ctx, command.DeleteProductCommand{
		Id: data.Id,
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
