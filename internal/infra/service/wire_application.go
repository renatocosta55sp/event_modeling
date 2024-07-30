package service

import (
	"context"
	"sync"

	"github.org/eventmodeling/product-management/internal/app"
	"github.org/eventmodeling/product-management/internal/app/command"
	"github.org/eventmodeling/product-management/internal/app/eventhandler"
	"github.org/eventmodeling/product-management/internal/app/query"
	"github.org/eventmodeling/product-management/internal/domain/product"
	"github.org/eventmodeling/product-management/internal/domain/product/events"
	appBB "github.org/eventmodeling/product-management/pkg/building_blocks/app"
	"github.org/eventmodeling/product-management/pkg/building_blocks/domain"
	"github.org/eventmodeling/product-management/pkg/building_blocks/infra/bus"
)

func NewApplication(ctx context.Context, wg *sync.WaitGroup, eventBus *bus.EventBus, repo product.ProductRepository) app.Application {

	eventChan := make(chan domain.Event, 1)

	eventBus.Subscribe(events.ProductCreatedEvent, eventChan)
	eventBus.Subscribe(events.ProductUpdatedEvent, eventChan)
	eventBus.Subscribe(events.ProductDeletedEvent, eventChan)

	eventHandlers := []appBB.EventHandlers{
		{
			EventName: events.ProductCreatedEvent,
			Handlers: []appBB.EventHandleable{
				eventhandler.NewCreateProductEventHandler(repo),
			},
			WgEnabled: true,
		},
		{
			EventName: events.ProductUpdatedEvent,
			Handlers: []appBB.EventHandleable{
				eventhandler.NewUpdateProductEventHandler(repo),
			},
			WgEnabled: true,
		},
		{
			EventName: events.ProductDeletedEvent,
			Handlers: []appBB.EventHandleable{
				eventhandler.NewDeleteProductEventHandler(repo),
			},
			WgEnabled: true,
		},
	}

	eventListener := appBB.NewEventListener(eventHandlers, eventBus, wg)
	go eventListener.Listen(ctx, eventChan)

	eventPublisher := bus.NewEventPublisher(eventBus)

	return app.Application{
		Commands: app.Commands{
			CreateProduct: command.NewProductHandler(eventPublisher),
			DeleteProduct: command.NewDeleteProductHandler(eventPublisher, repo),
			UpdateProduct: command.NewUpdateProductHandler(eventPublisher, repo),
		},
		Queries: app.Queries{
			AvailableProducts: query.NewAvailableProductsHandler(repo),
		},
	}
}
