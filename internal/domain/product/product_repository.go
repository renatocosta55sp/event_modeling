package product

import (
	"context"
)

type ProductRepository interface {
	Add(entity *ProductEntity, ctx context.Context) (int, error)
	Update(entity *ProductEntity, ctx context.Context) error
	GetAll(ctx context.Context) (*[]ProductEntity, error)
	GetById(id string, ctx context.Context) (*ProductEntity, error)
	Remove(entity *ProductEntity, ctx context.Context) error
}
