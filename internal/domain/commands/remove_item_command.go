package commands

import "github.com/google/uuid"

type RemoveItemCommand struct {
	AggregateID uuid.UUID
	ItemID      uuid.UUID
}
