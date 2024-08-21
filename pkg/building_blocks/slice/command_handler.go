package slice

import (
	"context"
	"errors"

	"github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/infra/bus"
)

type GenericCommandHandler struct {
}

func (GenericCommandHandler) Handle(ctxCancFunc context.CancelFunc, eventBus *bus.EventBus, eventResultChan chan bus.EventResult, domainEvent []domain.Event) (err error) {

	evPublisher := bus.NewEventPublisher(eventBus)
	evPublisher.Publish(domainEvent)

	eventResult, resultChanOk := <-eventResultChan

	if !resultChanOk {
		err = errors.New("result channel closed")
		return
	}

	if eventResult.Err != nil {
		err = eventResult.Err
		ctxCancFunc()
		return
	}

	return

}
