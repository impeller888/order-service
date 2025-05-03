package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusSubmitted  OrderStatus = "submitted"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type OrderStatus string

type Order struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Status    OrderStatus
	TotalCost uint
	Items     []*OrderItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItem struct {
	OrderID     uuid.UUID
	ProductID   uuid.UUID
	Description string
	Price       uint
	Quantity    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderRequest struct {
	UserID    uuid.UUID
	Status    *OrderStatus
	TotalCost *uint
	Items     []*OrderItemRequest
}

type OrderItemRequest struct {
	ProductID   uuid.UUID
	Description *string
	Price       *uint
	Quantity    *uint
}
