package app

import (
	"context"
	"sync"

	"github.org/eventmodeling/product-management/pkg/building_blocks/domain"
	"github.org/eventmodeling/product-management/pkg/building_blocks/infra/bus"
)

type EventHandlers struct {
	EventName string
	Handlers  []EventHandleable
	WgEnabled bool
}

type EventListener struct {
	eventHandlers []EventHandlers
	eventBus      *bus.EventBus
	wg            *sync.WaitGroup
}

func NewEventListener(eventHandlers []EventHandlers, eventBus *bus.EventBus, wg *sync.WaitGroup) *EventListener {
	return &EventListener{
		eventHandlers: eventHandlers,
		eventBus:      eventBus,
		wg:            wg,
	}
}

func (el *EventListener) Listen(ctx context.Context, eventChan <-chan domain.Event) {

	for {
		select {
		case event := <-eventChan:

			for _, handler := range el.eventHandlers {

				if handler.EventName == event.Type {
					el.dispatchHandlers(ctx, event, handler)
				}
			}

		}

	}

}

// handleHandlers dispatches the event to the event handlers
func (el *EventListener) dispatchHandlers(ctx context.Context, event domain.Event, handler EventHandlers) {

	for _, handlerFunc := range handler.Handlers {
		if err := handlerFunc.Handle(ctx, event); err != nil {
			// Remove the raised event on eventbus in case of an error
			el.eventBus.Remove(event.Type)
			el.eventBus.AddError(err)
		}
	}

	if handler.WgEnabled {
		el.wg.Done()
	}

}
