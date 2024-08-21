package testsuite

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.org/eventmodeling/ecommerce/internal/domain"

	"github.org/eventmodeling/ecommerce/internal/domain/commands"
	"github.org/eventmodeling/ecommerce/internal/events"
	"github.org/eventmodeling/ecommerce/internal/slices/additem"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/infra/bus"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/slice"
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

	err = slice.GenericCommandHandler{}.Handle(ctxCancFunc, eventBus, eventResultChan, cart.Events)

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
