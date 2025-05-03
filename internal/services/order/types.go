package order

import (
	"context"
	"local/order-service/internal/entity"

	"github.com/google/uuid"
)

type (
	orderRepository interface {
		GetOrderByID(ctx context.Context, orderID uuid.UUID) (*entity.Order, error)
		CreateOrder(ctx context.Context, req *entity.OrderRequest) (uuid.UUID, error)
	}

	userService interface {
		GetUserByID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	}

	productRepository interface {
		GetProductByID(ctx context.Context, productID uuid.UUID) (*entity.Product, error)
		UpdateProduct(ctx context.Context, productID uuid.UUID, entity *entity.ProductRequest) error
	}
)
