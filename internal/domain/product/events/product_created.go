package events

import "github.com/google/uuid"

const ProductCreatedEvent = "ProductCreatedEvent"

type ProductCreated struct {
	ID     uuid.UUID
	Entity any
}
