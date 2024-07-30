package query

import (
	"context"

	"github.org/napp/product-management/internal/domain/product"
	"github.org/napp/product-management/pkg/building_blocks/app"
)

type AvailableProducts struct {
}

type AvailableProductsHandler app.QueryHandler[AvailableProducts, *[]product.ProductEntity]

type availableProductsHandler struct {
	repo product.ProductRepository
}

func NewAvailableProductsHandler(repo product.ProductRepository) AvailableProductsHandler {
	return availableProductsHandler{repo: repo}
}

func (h availableProductsHandler) Handle(ctx context.Context, query AvailableProducts) (*[]product.ProductEntity, error) {
	entities, err := h.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return entities, nil
}
