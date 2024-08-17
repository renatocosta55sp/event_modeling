package events

import "github.com/google/uuid"

const ItemAddedEvent = "ItemAddedEvent"

type ItemAdded struct {
	AggregateId uuid.UUID
	Description string
	Price       float64
}
