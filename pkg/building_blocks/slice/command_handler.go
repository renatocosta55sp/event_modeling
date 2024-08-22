package slice

import (
	"context"
	"errors"

	"github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/infra/bus"
)

type GenericCommandHandler struct {
	CtxCancFunc     context.CancelFunc
	EventBus        *bus.EventBus
	EventResultChan chan bus.EventResult
}

func (g *GenericCommandHandler) Handle(domainEvent []domain.Event) (err error) {

	evPublisher := bus.NewEventPublisher(g.EventBus)
	evPublisher.Publish(domainEvent)

	eventResult, resultChanOk := <-g.EventResultChan

	if !resultChanOk {
		err = errors.New("result channel closed")
		return
	}

	if eventResult.Err != nil {
		err = eventResult.Err
		g.CtxCancFunc()
		return
	}

	return

}
