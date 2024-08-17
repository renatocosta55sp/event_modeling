package events

import "github.com/google/uuid"

const InventoryChangedEvent = "InventoryChangedEvent"

type InventoryChanged struct {
	ProductID uuid.UUID
	Inventory int
}
