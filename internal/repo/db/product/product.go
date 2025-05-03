package product

import (
	"local/order-service/internal/entity"
	"time"

	"github.com/google/uuid"
)

type product struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Quantity    uint      `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *product) fillFromEntity(entity *entity.Product) {
	p.ID = entity.ID
	p.Description = entity.Description
	p.Quantity = entity.Quantity
	p.CreatedAt = entity.CreatedAt
	p.UpdatedAt = entity.UpdatedAt
}

func (p *product) toEntity() *entity.Product {
	return &entity.Product{
		ID:          p.ID,
		Description: p.Description,
		Quantity:    p.Quantity,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
