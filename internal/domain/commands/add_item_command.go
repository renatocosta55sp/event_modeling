package commands

import "github.com/google/uuid"

type AddItemCommand struct {
	AggregateID uuid.UUID
	Description string
	Price       float64
}
