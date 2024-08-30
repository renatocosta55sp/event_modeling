package additem

import (
	"context"

	"github.com/renatocosta55sp/eventmodeling/internal/events"
	"github.com/renatocosta55sp/modeling/domain"
	"github.com/renatocosta55sp/modeling/infra/bus"
	"github.com/renatocosta55sp/modeling/slice"
)

func WireApp(ctx context.Context, eventBus *bus.EventBus) (eventRaisedChan chan bus.EventResult) {

	eventChan := make(chan domain.Event)

	eventBus.Subscribe(events.CartCreatedEvent, eventChan)
	eventBus.Subscribe(events.ItemAddedEvent, eventChan)

	eventHandlers := []slice.EventHandler{
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
