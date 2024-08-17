package additem

import (
	"context"

	"github.org/eventmodeling/ecommerce/internal/events"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/infra/bus"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/slice"
)

func WireApp(ctx context.Context, eventRaisedChan chan bus.EventRaised, eventBus *bus.EventBus) *bus.EventPublisher {

	eventChan := make(chan domain.Event, 1)

	eventBus.Subscribe(events.CartCreatedEvent, eventChan)
	eventBus.Subscribe(events.ItemAddedEvent, eventChan)

	eventHandlers := []slice.EventHandlers{
		{
			EventName: events.CartCreatedEvent,
			Handler:   NewCartEventHandler(),
		},
		{
			EventName: events.ItemAddedEvent,
			Handler:   NewItemEventHandler(),
			EndCycle:  true,
		},
	}

	eventListener := slice.NewEventListener(eventHandlers, eventBus, eventRaisedChan)
	go eventListener.Listen(ctx, eventChan)

	eventPublisher := bus.NewEventPublisher(eventBus)

	return eventPublisher
}
