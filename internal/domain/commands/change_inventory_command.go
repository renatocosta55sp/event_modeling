package commands

import "github.com/google/uuid"

type ChangeInventoryCommand struct {
	ProductID uuid.UUID
	Inventory int
}
