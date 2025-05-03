package response

import (
	"local/order-service/internal/entity"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID    `json:"id"`
	UserID    uuid.UUID    `json:"user_id"`
	Status    string       `json:"status"`
	TotalCost uint         `json:"total_cost"`
	Items     []*OrderItem `json:"items"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type OrderItem struct {
	ProductID   uuid.UUID `json:"product_id"`
	Description string    `json:"description"`
	Price       uint      `json:"price"`
	Quantity    uint      `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (o *Order) FillFromEntity(entity *entity.Order) {
	o.ID = entity.ID
	o.UserID = entity.UserID
	o.Status = string(entity.Status)
	o.TotalCost = entity.TotalCost
	o.CreatedAt = entity.CreatedAt
	o.UpdatedAt = entity.UpdatedAt

	o.Items = make([]*OrderItem, 0, len(entity.Items))

	for _, item := range entity.Items {
		orderItem := &OrderItem{}
		orderItem.fillFromEntity(item)
		o.Items = append(o.Items, orderItem)
	}
}

func (o *OrderItem) fillFromEntity(entity *entity.OrderItem) {
	o.ProductID = entity.ProductID
	o.Description = entity.Description
	o.Price = entity.Price
	o.Quantity = entity.Quantity
	o.CreatedAt = entity.CreatedAt
	o.UpdatedAt = entity.UpdatedAt
}
