package testsuite

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.org/eventmodeling/product-management/internal/app"
	"github.org/eventmodeling/product-management/internal/app/command"
	"github.org/eventmodeling/product-management/internal/domain/product/events"
	"github.org/eventmodeling/product-management/internal/infra/adapters/persistence"
	"github.org/eventmodeling/product-management/internal/infra/service"
	"github.org/eventmodeling/product-management/pkg/building_blocks/infra/bus"
)

var ag = &bus.AggregateRootTestCase{}
var eventBus = bus.NewEventBus()
var wireApp app.Application
var ctx context.Context
var wg sync.WaitGroup
var pgContainer testcontainers.Container

func init() {

	ctx = context.Background() // Initialize context here

	dbConn, container, err := InitTestContainer()
	if err != nil {
		log.Fatalf("Failed to initialize test container: %v", err)
	}
	pgContainer = container

	wireApp = service.NewApplication(ctx,
		&wg,
		eventBus,
		persistence.NewProductRepository(dbConn, "public"),
	)

}

func runCommand() {

	wg.Add(1)

	currentDateTime := time.Now().Format("2006-01-02 15:04:05.000")

	_, err := wireApp.Commands.CreateProduct.Handle(ctx, command.CreateProductCommand{
		Name:       "Product 0102023",
		Code:       123,
		Stock:      10,
		TotalStock: 10,
		CutStock:   10,
		PriceFrom:  10.00,
		PriceTo:    10.00,
		CreatedBy:  1,
		UpdatedBy:  1,
		CreatedAt:  currentDateTime,
		UpdatedAt:  currentDateTime,
	})
	if err != nil {
		ag.T.Fatal(err)
	}

	wg.Wait()
	if !eventBus.Raised(events.ProductCreatedEvent) {
		err := eventBus.GetError()
		ag.T.Fatal(err)
	}

}

func TestShouldBeAbleToRunEndToEndCommandSuccessfully(t *testing.T) {

	// Clean up the pg container
	defer func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	ag.T = t

	ag.
		Given(runCommand).
		When(eventBus.RaisedEvents()).
		Then(
			events.ProductCreatedEvent,
		).
		Assert()

}
