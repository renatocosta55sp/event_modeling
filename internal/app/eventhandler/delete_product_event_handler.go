package eventhandler

import (
	"context"

	"github.org/napp/product-management/internal/domain/product"
	"github.org/napp/product-management/internal/domain/product/events"
	"github.org/napp/product-management/pkg/building_blocks/app"
	"github.org/napp/product-management/pkg/building_blocks/domain"
)

type DeleteProductEventHandler struct {
	repo product.ProductRepository
}

func NewDeleteProductEventHandler(repo product.ProductRepository) app.EventHandleable {
	return &DeleteProductEventHandler{
		repo: repo,
	}
}

func (d DeleteProductEventHandler) Handle(ctx context.Context, event domain.Event) error {
	entity := event.Data.(events.ProductDeleted).Entity.(*product.ProductEntity)

	err := d.repo.Remove(entity, ctx)
	if err != nil {
		return err
	}

	return nil
}
