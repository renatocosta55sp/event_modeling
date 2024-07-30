package product

import (
	"context"
)

type ProductRepository interface {
	Add(entity *ProductEntity, ctx context.Context) (int, error)
	Update(entity *ProductEntity, ctx context.Context) error
	GetAll(ctx context.Context) (*[]ProductEntity, error)
	GetByCode(code int, ctx context.Context) (*ProductEntity, error)
	Remove(entity *ProductEntity, ctx context.Context) error
}
