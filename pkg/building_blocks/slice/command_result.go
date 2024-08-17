package slice

import "github.com/google/uuid"

type CommandResult struct {
	Identifier        uuid.UUID
	AggregateSequence int8
}
