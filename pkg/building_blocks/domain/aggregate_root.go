package domain

import (
	"time"

	"github.com/google/uuid"
)

type AggregateRoot struct {
	AggregateID uuid.NullUUID
	Version     int8
	Events      []Event `json:"-"`
}

func (a *AggregateRoot) RecordThat(event Event) {
	event.Timestamp = time.Now()
	a.Events = append(a.Events, event)
}
