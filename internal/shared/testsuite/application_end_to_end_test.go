package testsuite

import (
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"github.org/eventmodeling/ecommerce/internal/domain"

	"github.org/eventmodeling/ecommerce/internal/domain/commands"
	"github.org/eventmodeling/ecommerce/internal/events"
	"github.org/eventmodeling/ecommerce/internal/slices/additem"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/infra/bus"
)

var ag = &bus.AggregateRootTestCase{}
var eventBus = bus.NewEventBus()
var eventPublisher *bus.EventPublisher
var ctx context.Context
var eventRaisedChan chan bus.EventRaised

func init() {

	ctx = context.Background()
	//errors.Factory{}.Start()
	eventRaisedChan = make(chan bus.EventRaised)

	eventPublisher = additem.WireApp(ctx,
		eventRaisedChan,
		eventBus,
	)
}

func runCommand() {

	command := commands.AddItemCommand{
		Description: "Product 0102023",
		Price:       10.00,
	}
	cartItems := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}
	cart := domain.NewCart(uuid.NullUUID{}, cartItems, 1)

	_, err := cart.Handle(command)
	if err != nil {
		log.Fatal(err)
	}

	eventPublisher.Publish(cart.Events)

	eventRaised := <-eventRaisedChan

	if eventRaised.Err != nil {
		ag.T.Fatal(eventRaised.Err)
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
