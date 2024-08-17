package bus

import "github.org/eventmodeling/ecommerce/pkg/building_blocks/domain"

type EventRaised struct {
	Event domain.Event
	Err   error
}
