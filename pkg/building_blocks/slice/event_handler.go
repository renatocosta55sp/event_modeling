package slice

import (
	"context"

	"github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"
)

type EventHandleable interface {
	Handle(ctx context.Context, event domain.Event) error
}
