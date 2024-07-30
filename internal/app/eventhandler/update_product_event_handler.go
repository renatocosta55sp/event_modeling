package eventhandler

import (
	"context"

	"github.org/napp/product-management/internal/domain/product"
	"github.org/napp/product-management/internal/domain/product/events"
	"github.org/napp/product-management/pkg/building_blocks/app"
	"github.org/napp/product-management/pkg/building_blocks/domain"
)

type UpdateProductEventHandler struct {
	repo product.ProductRepository
}

func NewUpdateProductEventHandler(repo product.ProductRepository) app.EventHandleable {
	return &UpdateProductEventHandler{
		repo: repo,
	}
}

func (c UpdateProductEventHandler) Handle(ctx context.Context, event domain.Event) error {
	entity := event.Data.(events.ProductUpdated).Entity.(*product.ProductEntity)
	err := c.repo.Update(entity, ctx)
	if err != nil {
		return err
	}

	return nil
}
