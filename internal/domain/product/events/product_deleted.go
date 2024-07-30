package events

import "github.com/google/uuid"

const ProductDeletedEvent = "ProductDeletedEvent"

type ProductDeleted struct {
	ID     uuid.UUID
	Entity any
}
