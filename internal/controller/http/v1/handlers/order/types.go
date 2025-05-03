package order

import (
	"context"
	"local/order-service/internal/entity"

	"github.com/google/uuid"
)

type orderService interface {
	GetOrderByID(ctx context.Context, id uuid.UUID) (*entity.Order, error)
	CreateOrder(ctx context.Context, req *entity.OrderRequest) (uuid.UUID, error)
}
