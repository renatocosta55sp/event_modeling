package additem

import (
	"context"

	"github.com/renatocosta55sp/modeling/domain"
	"github.com/renatocosta55sp/modeling/slice"
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
