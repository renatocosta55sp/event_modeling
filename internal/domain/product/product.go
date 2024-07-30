package product

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.org/eventmodeling/product-management/internal/domain/product/events"
	"github.org/eventmodeling/product-management/pkg/building_blocks/domain"
	"github.org/eventmodeling/product-management/pkg/support"
)

type Product interface {
	isPriceFromLowerThanPriceTo() bool
	UpdateAvailableStock() int
	Create() Product
	Update() Product
	Delete() Product
	GetEvents() []domain.Event
}

type ProductEntity struct {
	domain.AggregateRoot
	Code           int
	Name           string
	Stock          int
	TotalStock     int
	CutStock       int
	AvailableStock int
	PriceFrom      float64
	PriceTo        float64
	CreatedAt      support.DateTime
	UpdatedAt      support.DateTime
	CreatedBy      int
	UpdatedBy      int
}

func NewProduct(
	aggregateId uuid.UUID,
	code int,
	name string,
	stock int,
	totalStock int,
	cutStock int,
	priceFrom float64,
	priceTo float64,
	createdAt support.DateTime,
	updatedAt support.DateTime,
	createdBy int,
	updatedBy int,
) (Product, error) {

	if name == "" {
		return nil, ErrEmptyName
	}

	productEntity := &ProductEntity{
		AggregateRoot: domain.AggregateRoot{ID: aggregateId},
		Code:          code,
		Name:          name,
		Stock:         stock,
		TotalStock:    totalStock,
		CutStock:      cutStock,
		PriceFrom:     priceFrom,
		PriceTo:       priceTo,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		CreatedBy:     createdBy,
		UpdatedBy:     updatedBy,
	}

	if productEntity.isPriceFromLowerThanPriceTo() {
		return nil, ErrPriceFromLowerThanPriceTo
	}

	return productEntity, nil
}

func (p *ProductEntity) UpdateAvailableStock() int {
	p.AvailableStock = p.TotalStock - p.CutStock
	return p.AvailableStock
}
func (p *ProductEntity) isPriceFromLowerThanPriceTo() bool {
	return p.PriceFrom < p.PriceTo
}

func (p *ProductEntity) Create() Product {

	// Raise event
	event := domain.Event{
		Type:      events.ProductCreatedEvent,
		Timestamp: time.Now(),
		Data: events.ProductCreated{
			ID:     p.ID,
			Entity: p,
		},
	}

	p.AggregateRoot.RecordThat(event)

	return p
}

func (p *ProductEntity) Update() Product {

	// Raise event
	event := domain.Event{
		Type:      events.ProductUpdatedEvent,
		Timestamp: time.Now(),
		Data: events.ProductUpdated{
			ID:     p.ID,
			Entity: p,
		},
	}

	p.AggregateRoot.RecordThat(event)

	return p
}

func (p *ProductEntity) Delete() Product {

	// Raise event
	event := domain.Event{
		Type:      events.ProductDeletedEvent,
		Timestamp: time.Now(),
		Data: events.ProductDeleted{
			ID:     p.ID,
			Entity: p,
		},
	}

	p.AggregateRoot.RecordThat(event)

	return p
}

func (p *ProductEntity) GetEvents() []domain.Event {
	return p.AggregateRoot.Events
}

var (
	ErrEmptyName                  = errors.New("product name is empty")
	ErrPriceFromLowerThanPriceTo  = errors.New("price from must be equal or greater than price to")
	ErrDifferentCreatorAndUpdater = errors.New("createdBy and UpdatedBy are different")
)
