package product

import (
	"context"
	"local/order-service/internal/entity"

	"github.com/google/uuid"
)

type productService interface {
	GetProductByID(ctx context.Context, id uuid.UUID) (*entity.Product, error)
	CreateProduct(ctx context.Context, product *entity.ProductRequest) (uuid.UUID, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, product *entity.ProductRequest) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}
