package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.org/napp/product-management/pkg/support"
)

func TestInvalidArguments(t *testing.T) {

	createdAt, err := support.ParseDateTime("2020-01-01 00:00:00.000")
	assert.NoError(t, err)
	updatedAt, err := support.ParseDateTime("2020-01-01 00:00:00.000")
	assert.NoError(t, err)

	var tests = []struct {
		Code           int
		Name           string
		Stock          int
		TotalStock     int
		CutStock       int
		AvailableStock int
		PriceFrom      float64
		PriceTo        float64
		CreatedAt      string
		UpdatedAt      string
		CreatedBy      int
		UpdatedBy      int
		want           error
	}{
		{

			Code:           123,
			Name:           "Product 01",
			Stock:          10,
			TotalStock:     10,
			CutStock:       10,
			AvailableStock: 10,
			PriceFrom:      9,
			PriceTo:        10,
			CreatedAt:      createdAt.String(),
			UpdatedAt:      updatedAt.String(),
			CreatedBy:      1,
			UpdatedBy:      1,
			want:           ErrPriceFromLowerThanPriceTo,
		},
		{

			Code:           123,
			Name:           "Product 02",
			Stock:          10,
			TotalStock:     10,
			CutStock:       10,
			AvailableStock: 10,
			PriceFrom:      8,
			PriceTo:        10,
			CreatedAt:      createdAt.String(),
			UpdatedAt:      updatedAt.String(),
			CreatedBy:      1,
			UpdatedBy:      1,
			want:           ErrPriceFromLowerThanPriceTo,
		},
		{

			Code:           123,
			Name:           "",
			Stock:          10,
			TotalStock:     10,
			CutStock:       10,
			AvailableStock: 10,
			PriceFrom:      10,
			PriceTo:        10,
			CreatedAt:      createdAt.String(),
			UpdatedAt:      updatedAt.String(),
			CreatedBy:      1,
			UpdatedBy:      1,
			want:           ErrEmptyName,
		},
	}

	for _, test := range tests {
		_, err := NewProduct(
			test.Code,
			test.Name,
			test.Stock,
			test.TotalStock,
			test.CutStock,
			test.PriceFrom,
			test.PriceTo,
			createdAt,
			updatedAt,
			test.CreatedBy,
			test.UpdatedBy,
		)
		assert.Equal(t, err, test.want, "Expected: %d - Got: %d", test.want, err)
	}

}
