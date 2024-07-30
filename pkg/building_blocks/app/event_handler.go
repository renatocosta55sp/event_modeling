package app

import (
	"context"

	"github.org/eventmodeling/product-management/pkg/building_blocks/domain"
)

type EventHandleable interface {
	Handle(ctx context.Context, event domain.Event) error
}
