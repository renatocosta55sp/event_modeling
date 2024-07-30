package command

import (
	"context"

	"github.org/napp/product-management/internal/domain/product"
	"github.org/napp/product-management/pkg/building_blocks/app"

	"github.org/napp/product-management/pkg/building_blocks/infra/bus"
)

type DeleteProductCommand struct {
	Code int
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

	productRead, err := h.repo.GetByCode(cmd.Code, ctx)

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
