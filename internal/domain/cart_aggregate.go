package domain

import (
	"errors"
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/renatocosta55sp/eventmodeling/internal/domain/commands"
	"github.com/renatocosta55sp/eventmodeling/internal/events"
	"github.com/renatocosta55sp/modeling/domain"
	"github.com/renatocosta55sp/modeling/slice"
)

type CartAggregate struct {
	domain.AggregateRoot
	CartItems []uuid.UUID
	CreatedBy int16
}

func NewCart(
	aggregateId uuid.NullUUID,
	cartItems []uuid.UUID,
	createdBy int16,
) *CartAggregate {

	cart := &CartAggregate{
		AggregateRoot: domain.AggregateRoot{AggregateID: aggregateId, Version: CartAggregateVersion},
		CartItems:     cartItems,
		CreatedBy:     createdBy,
	}

	return cart
}

func (c *CartAggregate) isEligiblePrice(price float64) bool {
	return price > 0
}

func (c *CartAggregate) Handle(command commands.AddItemCommand) (slice.CommandResult, error) {

	if !c.AggregateID.Valid {

		c.AggregateRoot.RecordThat(
			domain.Event{
				Type: events.CartCreatedEvent,
				Data: events.CartCreated{
					AggregateId: command.AggregateID,
				},
			},
		)
	}

	if !c.isEligiblePrice(command.Price) {
		return slice.CommandResult{
			Identifier:        command.AggregateID,
			AggregateSequence: CartAggregateVersion,
		}, ErrPriceLess
	}

	if len(c.CartItems) > 3 {
		return slice.CommandResult{
			Identifier:        command.AggregateID,
			AggregateSequence: CartAggregateVersion,
		}, ErrItemsExceeded
	}

	c.AggregateRoot.RecordThat(
		domain.Event{
			Type: events.ItemAddedEvent,
			Data: events.ItemAdded{
				AggregateId: command.AggregateID,
				Description: command.Description,
				Price:       command.Price,
			},
		},
	)

	return slice.CommandResult{
		Identifier:        command.AggregateID,
		AggregateSequence: CartAggregateVersion,
	}, nil

}

func (c *CartAggregate) HandleRemoveItem(command commands.RemoveItemCommand) (slice.CommandResult, error) {

	if !slices.Contains(c.CartItems, command.ItemID) {
		return slice.CommandResult{
			Identifier:        command.AggregateID,
			AggregateSequence: CartAggregateVersion,
		}, fmt.Errorf(ErrItemNotFound.Error(), command.ItemID)

	}

	c.AggregateRoot.RecordThat(
		domain.Event{
			Type: events.ItemRemovedEvent,
			Data: events.ItemRemoved{
				AggregateId: command.AggregateID,
				ItemId:      command.ItemID,
			},
		},
	)

	return slice.CommandResult{
		Identifier:        command.AggregateID,
		AggregateSequence: CartAggregateVersion,
	}, nil
}

var CartAggregateVersion = int8(1)

var (
	ErrEmptyName     = errors.New("error.cart.name.required")
	ErrPriceLess     = errors.New("error.cart.price.priceless")
	ErrItemsExceeded = errors.New("error.cart.cartitems.exceeded")
	ErrItemNotFound  = errors.New("error.cart.item.notfound")
)
