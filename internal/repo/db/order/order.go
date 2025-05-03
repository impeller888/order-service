package order

import (
	"local/order-service/internal/entity"
	"time"

	"github.com/google/uuid"
)

type order struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Status    string
	TotalCost uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type orderItem struct {
	OrderID     uuid.UUID
	ProductID   uuid.UUID
	Description string
	Price       uint
	Quantity    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (o *order) fillFromEntity(entity *entity.Order) {
	o.ID = entity.ID
	o.UserID = entity.UserID
	o.Status = string(entity.Status)
	o.TotalCost = entity.TotalCost
	o.CreatedAt = entity.CreatedAt
	o.UpdatedAt = entity.UpdatedAt
}

func (o *order) toEntity() *entity.Order {
	return &entity.Order{
		ID:        o.ID,
		UserID:    o.UserID,
		Status:    entity.OrderStatus(o.Status),
		TotalCost: o.TotalCost,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

func (o *orderItem) toEntity() *entity.OrderItem {
	return &entity.OrderItem{
		OrderID:     o.OrderID,
		ProductID:   o.ProductID,
		Description: o.Description,
		Price:       o.Price,
		Quantity:    o.Quantity,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
	}
}

func (o *orderItem) fillFromEntity(entity *entity.OrderItem) {
	o.OrderID = entity.OrderID
	o.ProductID = entity.ProductID
	o.Description = entity.Description
	o.Price = entity.Price
	o.Quantity = entity.Quantity
	o.CreatedAt = entity.CreatedAt
	o.UpdatedAt = entity.UpdatedAt
}
