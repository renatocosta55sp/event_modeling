package additem

import (
	"context"

	"github.org/eventmodeling/ecommerce/internal/events"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/infra/bus"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/slice"
)

func WireApp(ctx context.Context, eventBus *bus.EventBus) (eventRaisedChan chan bus.EventResult) {

	eventChan := make(chan domain.Event)

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

	eventRaisedChan = make(chan bus.EventResult)

	eventListener := slice.NewEventListener(eventHandlers, eventBus, eventRaisedChan)
	go eventListener.Listen(ctx, eventChan)

	return
}
