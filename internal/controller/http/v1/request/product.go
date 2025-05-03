package request

import "local/order-service/internal/entity"

type ProductRequest struct {
	Description *string `json:"description"`
	Quantity    *uint   `json:"quantity"`
}

func (p *ProductRequest) ToEntity() *entity.ProductRequest {
	return &entity.ProductRequest{
		Description: p.Description,
		Quantity:    p.Quantity,
	}
}
