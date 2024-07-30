package events

import "github.com/google/uuid"

const ProductUpdatedEvent = "ProductUpdatedEvent"

type ProductUpdated struct {
	ID     uuid.UUID
	Entity any
}
