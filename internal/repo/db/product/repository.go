package product

import (
	"context"
	"errors"
	"fmt"
	"local/order-service/internal/entity"
	"local/order-service/internal/repo/db"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	raw := &product{}
	err := r.db.QueryRow(ctx, "SELECT * FROM products WHERE id = $1", id).Scan(
		&raw.ID, &raw.Description, &raw.Quantity, &raw.CreatedAt, &raw.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, entity.ErrNotFound
		}
		return nil, err
	}

	return raw.toEntity(), nil

}

func (r *ProductRepository) CreateProduct(ctx context.Context, req *entity.ProductRequest) (uuid.UUID, error) {
	if req.Description == nil {
		return uuid.Nil, db.ErrMissedRequiredField
	}
	if req.Quantity == nil {
		val := uint(0)
		req.Quantity = &val
	}

	var id uuid.UUID
	err := r.db.QueryRow(ctx, "INSERT INTO products (description, quantity) VALUES ($1, $2) RETURNING id", *req.Description, *req.Quantity).
		Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id uuid.UUID, req *entity.ProductRequest) error {

	var (
		namedArgs = pgx.NamedArgs{
			"id": id,
		}
		cols = []string{}
	)

	if req.Description != nil {
		namedArgs["desc"] = *req.Description
		cols = append(cols, "description = @desc")
	}
	if req.Quantity != nil {
		namedArgs["quantity"] = *req.Quantity
		cols = append(cols, "quantity = @quantity")
	}

	if len(namedArgs) == 0 {
		return nil
	}

	_, err := r.db.Exec(ctx, fmt.Sprintf("UPDATE products SET %s WHERE id = @id", strings.Join(cols, ",")), namedArgs)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, "DELETE FROM products WHERE id = $1", id)
	return err
}
