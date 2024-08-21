package bus

import "github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"

type EventResult struct {
	Event domain.Event
	Err   error
}
