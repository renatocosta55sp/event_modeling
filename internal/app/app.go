package app

import (
	"github.org/eventmodeling/product-management/internal/app/command"
	"github.org/eventmodeling/product-management/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateProduct command.CreateProductHandler
	UpdateProduct command.UpdateProductHandler
	DeleteProduct command.DeleteProductHandler
}

type Queries struct {
	AvailableProducts query.AvailableProductsHandler
}
