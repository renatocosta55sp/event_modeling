package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"github.org/napp/product-management/internal/domain/product"
)

const ProductTableName = "products"

type RepoProduct struct {
	Conn     *pgx.Conn
	DBSchema string
}

func (r *RepoProduct) Add(entity *product.ProductEntity, ctx context.Context) (int, error) {

	query := fmt.Sprintf(`INSERT INTO %s.%s (aggregate_identifier, code, name, stock, total_stock, cut_stock, available_stock, price_from, price_to, created_at, updated_at, created_by, updated_by) VALUES (@aggregateIdentifier, @code, @name, @stock, @totalStock, @cutStock, @availableStock, @priceFrom, @priceTo, @createdAt, @updatedAt, @createdBy, @updatedBy) RETURNING code`, r.DBSchema, ProductTableName)

	args := pgx.NamedArgs{
		"aggregateIdentifier": entity.ID,
		"code":                entity.Code,
		"name":                entity.Name,
		"stock":               entity.Stock,
		"totalStock":          entity.TotalStock,
		"cutStock":            entity.CutStock,
		"availableStock":      entity.AvailableStock,
		"priceFrom":           entity.PriceFrom,
		"priceTo":             entity.PriceTo,
		"createdAt":           entity.CreatedAt,
		"updatedAt":           entity.UpdatedAt,
		"createdBy":           entity.CreatedBy,
		"updatedBy":           entity.UpdatedBy,
	}

	logrus.Info("args", args)

	var product_id int
	err := r.Conn.QueryRow(ctx, query, args).Scan(&product_id)

	if err != nil {
		return 0, fmt.Errorf("unable to insert row: %w", err)
	}
	return product_id, nil
}

func (r *RepoProduct) Update(entity *product.ProductEntity, ctx context.Context) error {

	command, err := r.Conn.Exec(ctx, "update products set name=$1, stock=$2, total_stock=$3, cut_stock=$4, available_stock=$5, price_from=$6, price_to=$7, updated_at=$8, updated_by=$9 where code=$10", entity.Name, entity.Stock, entity.TotalStock, entity.CutStock, entity.AvailableStock, entity.PriceFrom, entity.PriceTo, entity.UpdatedAt, entity.UpdatedBy, entity.Code)
	if err != nil {
		return err
	}

	if command.RowsAffected() != 1 {
		return errors.New("no row affected to update")
	}

	return nil
}

func (r *RepoProduct) GetAll(ctx context.Context) (*[]product.ProductEntity, error) {

	rows, err := r.Conn.Query(ctx, "SELECT aggregate_identifier, code, name, stock, total_stock, cut_stock, available_stock, price_from, price_to, created_at, updated_at, created_by, updated_by FROM products")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var products []product.ProductEntity
	for rows.Next() {
		var p product.ProductEntity
		if err := rows.Scan(&p.ID, &p.Code, &p.Name, &p.Stock, &p.TotalStock, &p.CutStock, &p.AvailableStock, &p.PriceFrom, &p.PriceTo, &p.CreatedAt, &p.UpdatedAt, &p.CreatedBy, &p.UpdatedBy); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return &products, nil
}

func (r *RepoProduct) Remove(entity *product.ProductEntity, ctx context.Context) error {
	command, err := r.Conn.Exec(ctx, "delete from products where code=$1", entity.Code)
	if err != nil {
		return err
	}
	if command.RowsAffected() != 1 {
		return errors.New("no row found to delete")
	}
	return nil
}

func (r *RepoProduct) GetByCode(code int, ctx context.Context) (*product.ProductEntity, error) {
	var p product.ProductEntity
	err := r.Conn.QueryRow(ctx, "select aggregate_identifier, code, name, stock, total_stock, cut_stock, available_stock, price_from, price_to, created_at, updated_at, created_by, updated_by from products where code=$1", code).Scan(&p.ID, &p.Code, &p.Name, &p.Stock, &p.TotalStock, &p.CutStock, &p.AvailableStock, &p.PriceFrom, &p.PriceTo, &p.CreatedAt, &p.UpdatedAt, &p.CreatedBy, &p.UpdatedBy)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func NewProductRepository(conn *pgx.Conn, dbSchema string) product.ProductRepository {
	return &RepoProduct{Conn: conn, DBSchema: dbSchema}
}
