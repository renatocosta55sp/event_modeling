package slice

import (
	"context"

	"github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/infra/bus"
)

type EventHandler struct {
	EventName string
	Handler   EventHandleable
	EndCycle  bool
}

type EventListener struct {
	eventHandlers   []EventHandler
	eventBus        *bus.EventBus
	eventResultChan chan bus.EventResult
}

func NewEventListener(eventHandlers []EventHandler, eventBus *bus.EventBus, eventsResultChan chan bus.EventResult) *EventListener {
	return &EventListener{
		eventHandlers:   eventHandlers,
		eventBus:        eventBus,
		eventResultChan: eventsResultChan,
	}
}

// Listen listens for incoming events and processes them
func (el *EventListener) Listen(ctx context.Context, eventChan <-chan domain.Event) {

	for {
		select {
		case event := <-eventChan:
			go el.dispatchToHandlers(ctx, event)
		case <-ctx.Done():
			close(el.eventResultChan)
			return
		}
	}

}

// dispatchToHandlers dispatches the event to the appropriate event handlers
func (el *EventListener) dispatchToHandlers(ctx context.Context, event domain.Event) {
	for _, handler := range el.eventHandlers {
		if handler.EventName == event.Type {

			go func(handler EventHandler) {

				err := handler.Handler.Handle(ctx, event)
				eventResult := bus.EventResult{Event: event}
				if err != nil {
					eventResult.Err = err
				}

				el.eventResultChan <- eventResult
				if handler.EndCycle {
					return
				}

			}(handler)

		}
	}
}
