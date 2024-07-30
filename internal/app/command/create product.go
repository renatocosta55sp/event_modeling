package command

import (
	"context"

	"github.com/google/uuid"
	"github.org/napp/product-management/internal/domain/product"
	"github.org/napp/product-management/pkg/building_blocks/app"

	"github.org/napp/product-management/pkg/building_blocks/infra/bus"
	"github.org/napp/product-management/pkg/support"
)

type CreateProductCommand struct {
	Code       int
	Name       string
	Stock      int
	TotalStock int
	CutStock   int
	PriceFrom  float64
	PriceTo    float64
	CreatedBy  int
	UpdatedBy  int
	CreatedAt  string
	UpdatedAt  string
}

type CreateProductHandler app.CommandHandler[CreateProductCommand, product.Product]

type createProductHandler struct {
	eventPublisher *bus.EventPublisher
}

func NewProductHandler(eventPublisher *bus.EventPublisher) CreateProductHandler {
	return createProductHandler{
		eventPublisher: eventPublisher,
	}
}

func (h createProductHandler) Handle(ctx context.Context, cmd CreateProductCommand) (product.Product, error) {

	createdAt, err := support.ParseDateTime(cmd.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := support.ParseDateTime(cmd.UpdatedAt)
	if err != nil {
		return nil, err
	}

	product, err := product.NewProduct(
		uuid.New(),
		cmd.Code,
		cmd.Name,
		cmd.Stock,
		cmd.TotalStock,
		cmd.CutStock,
		cmd.PriceFrom,
		cmd.PriceTo,
		createdAt,
		updatedAt,
		cmd.CreatedBy,
		cmd.UpdatedBy,
	)

	if err != nil {
		return nil, err
	}

	product.UpdateAvailableStock()
	product.Create()

	h.eventPublisher.PublishEvents(product.GetEvents())

	return product, nil
}
