package bus

import "github.org/eventmodeling/product-management/pkg/building_blocks/domain"

type EventPublisher struct {
	eventBus *EventBus
}

func NewEventPublisher(eventBus *EventBus) *EventPublisher {
	return &EventPublisher{
		eventBus: eventBus,
	}
}

func (ep *EventPublisher) PublishEvents(events []domain.Event) {

	for _, event := range events {
		ep.eventBus.Publish(event)
	}
}
