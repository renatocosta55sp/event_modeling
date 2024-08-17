package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.org/eventmodeling/ecommerce/internal/domain/commands"
)

func TestInvalidArguments(t *testing.T) {

	cartItems := []uuid.UUID{uuid.New(), uuid.New(), uuid.New(), uuid.New()}
	cart := NewCart(uuid.NullUUID{UUID: uuid.New(), Valid: true}, cartItems, 1)

	var tests = []struct {
		description string
		price       float64
		createdBy   int16
		want        error
	}{
		{

			description: "Product 01",
			price:       0,
			createdBy:   10,
			want:        ErrPriceLess,
		},
		{

			description: "Product 01",
			price:       10,
			createdBy:   10,
			want:        ErrItemsExceeded,
		},
	}

	for _, test := range tests {
		_, err := cart.Handle(
			commands.AddItemCommand{
				AggregateID: cart.AggregateID.UUID,
				Description: test.description,
				Price:       test.price,
			},
		)
		assert.Equal(t, err, test.want, "Expected: %d - Got: %d", test.want, err)
	}

}
