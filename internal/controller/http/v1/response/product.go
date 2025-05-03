package response

import (
	"local/order-service/internal/entity"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Quantity    uint      `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p *Product) FillFromEntity(entity *entity.Product) {
	p.ID = entity.ID
	p.Description = entity.Description
	p.Quantity = entity.Quantity
	p.CreatedAt = entity.CreatedAt
	p.UpdatedAt = entity.UpdatedAt
}
