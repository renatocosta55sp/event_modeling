package eventhandler

import (
	"context"

	"github.org/eventmodeling/product-management/internal/domain/product"
	"github.org/eventmodeling/product-management/internal/domain/product/events"
	"github.org/eventmodeling/product-management/pkg/building_blocks/app"
	"github.org/eventmodeling/product-management/pkg/building_blocks/domain"
)

type CreateProductEventHandler struct {
	repo product.ProductRepository
}

func NewCreateProductEventHandler(repo product.ProductRepository) app.EventHandleable {
	return &CreateProductEventHandler{
		repo: repo,
	}
}

func (c CreateProductEventHandler) Handle(ctx context.Context, event domain.Event) error {
	entity := event.Data.(events.ProductCreated).Entity.(*product.ProductEntity)
	_, err := c.repo.Add(entity, ctx)
	if err != nil {
		return err
	}

	return nil
}
