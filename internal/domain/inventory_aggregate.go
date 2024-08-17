package domain

import (
	"github.com/google/uuid"
	"github.org/eventmodeling/ecommerce/internal/domain/commands"
	"github.org/eventmodeling/ecommerce/internal/events"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"
	"github.org/eventmodeling/ecommerce/pkg/building_blocks/slice"
)

type InventoryAggregate struct {
	domain.AggregateRoot
	CreatedBy int16
}

func NewInventory(
	aggregateId uuid.NullUUID,
	createdBy int16,
) *InventoryAggregate {

	inventory := &InventoryAggregate{
		AggregateRoot: domain.AggregateRoot{AggregateID: aggregateId, Version: InventoryAggregateVersion},
		CreatedBy:     createdBy,
	}

	return inventory
}

func (i *InventoryAggregate) Handle(command commands.ChangeInventoryCommand) slice.CommandResult {

	i.AggregateRoot.RecordThat(
		domain.Event{
			Type: events.InventoryChangedEvent,
			Data: events.InventoryChanged{
				ProductID: command.ProductID,
				Inventory: command.Inventory,
			},
		},
	)

	return slice.CommandResult{
		Identifier:        command.ProductID,
		AggregateSequence: CartAggregateVersion,
	}

}

var InventoryAggregateVersion = int8(1)
