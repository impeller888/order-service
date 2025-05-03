package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	Description string
	Quantity    uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductRequest struct {
	Description *string
	Quantity    *uint
}
