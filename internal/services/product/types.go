package product

import (
	"context"
	"local/order-service/internal/entity"

	"github.com/google/uuid"
)

type productRepository interface {
	GetProductByID(ctx context.Context, id uuid.UUID) (*entity.Product, error)
	CreateProduct(ctx context.Context, req *entity.ProductRequest) (uuid.UUID, error)
	UpdateProduct(ctx context.Context, id uuid.UUID, req *entity.ProductRequest) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}
