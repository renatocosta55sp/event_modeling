package testsuite

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/renatocosta55sp/eventmodeling/internal/domain"
	"github.com/renatocosta55sp/eventmodeling/internal/domain/commands"
	"github.com/renatocosta55sp/eventmodeling/internal/events"
	"github.com/renatocosta55sp/eventmodeling/internal/slices/additem"
	"github.com/renatocosta55sp/modeling/infra/bus"
	"github.com/renatocosta55sp/modeling/slice"
	"github.com/stretchr/testify/assert"
)

var ag = &bus.AggregateRootTestCase{}
var eventBus = bus.NewEventBus()
var ctx context.Context
var ctxCancFunc context.CancelFunc
var eventResultChan chan bus.EventResult

func init() {

	ctx, ctxCancFunc = context.WithTimeout(context.Background(), 5*time.Second)
	//errors.Factory{}.Start()

	eventResultChan = additem.WireApp(ctx,
		eventBus,
	)
}

func runCommand() {

	aggregateIdentifier := uuid.New()
	command := commands.AddItemCommand{
		AggregateID: aggregateIdentifier,
		Description: "Product 0102023",
		Price:       10.00,
	}
	cartItems := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}
	cart := domain.NewCart(uuid.NullUUID{}, cartItems, 1)

	commandResult, err := cart.Handle(command)
	if err != nil {
		ag.T.Fatal(err)
	}

	commandResultToCompare := slice.CommandResult{
		Identifier:        aggregateIdentifier,
		AggregateSequence: domain.CartAggregateVersion,
	}

	assert.Equal(ag.T, commandResult, commandResultToCompare, "The CommandResult should be equal")

	err = (&slice.GenericCommandHandler{
		EventBus:        eventBus,
		CtxCancFunc:     ctxCancFunc,
		EventResultChan: eventResultChan,
	}).Handle(cart.Events)

	if err != nil {
		ag.T.Fatal(err)
	}

}

func TestShouldBeAbleToRunEndToEndCommandSuccessfully(t *testing.T) {

	ag.T = t

	ag.
		Given(runCommand).
		When(eventBus.RaisedEvents()).
		Then(
			events.CartCreatedEvent,
			events.ItemAddedEvent,
		).
		Assert()

}
