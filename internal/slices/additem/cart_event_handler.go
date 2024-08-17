package additem

import (
	"context"

	"github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/slice"
)

type CartEventHandler struct {
	//repo product.ProductRepository
}

func NewCartEventHandler() slice.EventHandleable {
	return &CartEventHandler{}
}

func (i CartEventHandler) Handle(ctx context.Context, event domain.Event) error {
	return nil
}

type ItemEventHandler struct {
	//repo product.ProductRepository
}

func NewItemEventHandler() slice.EventHandleable {
	return &ItemEventHandler{}
}

func (i ItemEventHandler) Handle(ctx context.Context, event domain.Event) error {
	return nil
}
