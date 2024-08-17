package events

import "github.com/google/uuid"

const ItemRemovedEvent = "ItemRemovedEvent"

type ItemRemoved struct {
	AggregateId uuid.UUID
	ItemId      uuid.UUID
}
