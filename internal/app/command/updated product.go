package command

import (
	"context"

	"github.org/eventmodeling/product-management/internal/domain/product"
	"github.org/eventmodeling/product-management/pkg/building_blocks/app"
	"github.org/eventmodeling/product-management/pkg/support"

	"github.org/eventmodeling/product-management/pkg/building_blocks/infra/bus"
)

type UpdateProductCommand struct {
	Id         string
	Name       string
	Stock      int
	TotalStock int
	CutStock   int
	PriceFrom  float64
	PriceTo    float64
	UpdatedBy  int
	UpdatedAt  string
}

type UpdateProductHandler app.CommandHandler[UpdateProductCommand, product.Product]

type updateProductHandler struct {
	eventPublisher *bus.EventPublisher
	repo           product.ProductRepository
}

func NewUpdateProductHandler(eventPublisher *bus.EventPublisher, repo product.ProductRepository) UpdateProductHandler {
	return updateProductHandler{
		eventPublisher: eventPublisher,
		repo:           repo,
	}
}

func (h updateProductHandler) Handle(ctx context.Context, cmd UpdateProductCommand) (product.Product, error) {

	productRead, err := h.repo.GetById(cmd.Id, ctx)

	if err != nil {
		return nil, err
	}

	updatedAt, err := support.ParseDateTime(cmd.UpdatedAt)
	if err != nil {
		return nil, err
	}

	product, err := product.NewProduct(
		productRead.ID,
		productRead.Code,
		cmd.Name,
		cmd.Stock,
		cmd.TotalStock,
		cmd.CutStock,
		cmd.PriceFrom,
		cmd.PriceTo,
		productRead.CreatedAt,
		updatedAt,
		productRead.CreatedBy,
		cmd.UpdatedBy,
	)

	if err != nil {
		return nil, err
	}

	product.UpdateAvailableStock()
	product.Update()

	h.eventPublisher.PublishEvents(product.GetEvents())

	return nil, nil
}
