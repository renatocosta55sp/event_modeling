package slice

import (
	"context"

	"github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/infra/bus"
)

type EventHandlers struct {
	EventName string
	Handler   EventHandleable
	EndCycle  bool
}

type EventListener struct {
	eventHandlers   []EventHandlers
	eventBus        *bus.EventBus
	eventRaisedChan chan bus.EventResult
}

func NewEventListener(eventHandlers []EventHandlers, eventBus *bus.EventBus, eventsRaisedChan chan bus.EventResult) *EventListener {
	return &EventListener{
		eventHandlers:   eventHandlers,
		eventBus:        eventBus,
		eventRaisedChan: eventsRaisedChan,
	}
}

func (el *EventListener) Listen(ctx context.Context, eventChan <-chan domain.Event) {

	for {
		select {
		case event := <-eventChan:
			el.dispatchToHandlers(ctx, event)
		case <-ctx.Done():
			// Signal completion
			close(el.eventRaisedChan)
			return
		}
	}

}

// handleHandlers dispatches the event to the event handlers
func (el *EventListener) dispatchToHandlers(ctx context.Context, event domain.Event) {
	for _, evh := range el.eventHandlers {
		if evh.EventName == event.Type {
			err := evh.Handler.Handle(ctx, event)
			eventRaised := bus.EventResult{
				Event: event,
			}

			if err != nil {
				eventRaised.Err = err
				el.eventRaisedChan <- eventRaised
				return
			}

			if evh.EndCycle {
				eventRaised.Event = event
				el.eventRaisedChan <- eventRaised
				return
			}

		}
	}

}
