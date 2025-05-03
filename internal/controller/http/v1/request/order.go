package request

import (
	"local/order-service/internal/entity"

	"github.com/google/uuid"
)

type OrderRequest struct {
	UserID    uuid.UUID           `json:"user_id"`
	Status    *string             `json:"status"`
	TotalCost *uint               `json:"total_cost"`
	Items     []*OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductID   uuid.UUID `json:"product_id"`
	Description *string   `json:"description"`
	Price       *uint     `json:"price"`
	Quantity    *uint     `json:"quantity"`
}

func (o *OrderRequest) ToEntity() *entity.OrderRequest {
	req := &entity.OrderRequest{
		UserID:    o.UserID,
		TotalCost: o.TotalCost,
	}
	if o.Status != nil {
		status := *o.Status
		orderStatus := entity.OrderStatus(status)
		req.Status = &orderStatus
	}

	req.Items = make([]*entity.OrderItemRequest, 0, len(o.Items))

	for _, item := range o.Items {
		req.Items = append(req.Items, item.ToEntity())
	}

	return req
}

func (o *OrderItemRequest) ToEntity() *entity.OrderItemRequest {
	return &entity.OrderItemRequest{
		ProductID:   o.ProductID,
		Description: o.Description,
		Price:       o.Price,
		Quantity:    o.Quantity,
	}
}
