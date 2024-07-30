package command

import (
	"context"

	"github.org/eventmodeling/product-management/internal/domain/product"
	"github.org/eventmodeling/product-management/pkg/building_blocks/app"

	"github.org/eventmodeling/product-management/pkg/building_blocks/infra/bus"
)

type DeleteProductCommand struct {
	Id string
}

type DeleteProductHandler app.CommandHandler[DeleteProductCommand, product.Product]

type deleteProductHandler struct {
	eventPublisher *bus.EventPublisher
	repo           product.ProductRepository
}

func NewDeleteProductHandler(eventPublisher *bus.EventPublisher, repo product.ProductRepository) DeleteProductHandler {
	return deleteProductHandler{
		eventPublisher: eventPublisher,
		repo:           repo,
	}
}

func (h deleteProductHandler) Handle(ctx context.Context, cmd DeleteProductCommand) (product.Product, error) {

	productRead, err := h.repo.GetById(cmd.Id, ctx)

	if err != nil {
		return nil, err
	}

	product, err := product.NewProduct(
		productRead.ID,
		productRead.Code,
		productRead.Name,
		productRead.Stock,
		productRead.TotalStock,
		productRead.CutStock,
		productRead.PriceFrom,
		productRead.PriceTo,
		productRead.CreatedAt,
		productRead.UpdatedAt,
		productRead.CreatedBy,
		productRead.UpdatedBy,
	)

	if err != nil {
		return nil, err
	}

	product.Delete()

	h.eventPublisher.PublishEvents(product.GetEvents())

	return nil, nil
}
