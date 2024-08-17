package events

import "github.com/google/uuid"

const CartCreatedEvent = "CartCreatedEvent"

type CartCreated struct {
	AggregateId uuid.UUID
}
