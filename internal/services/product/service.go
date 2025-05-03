package product

import (
	"context"
	"local/order-service/internal/entity"

	"github.com/google/uuid"
)

type ProductService struct {
	repo productRepository
}

func NewProductService(repo productRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProductByID(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	return s.repo.GetProductByID(ctx, id)
}

func (s *ProductService) CreateProduct(ctx context.Context, req *entity.ProductRequest) (uuid.UUID, error) {
	if req.Quantity == nil {
		return uuid.Nil, &entity.ErrFieldRequired{
			Name: "quantity",
		}
	}
	if req.Description == nil {
		return uuid.Nil, &entity.ErrFieldRequired{
			Name: "description",
		}
	}

	return s.repo.CreateProduct(ctx, req)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id uuid.UUID, req *entity.ProductRequest) error {
	_, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return err
	}

	if req.Quantity != nil && *req.Quantity < 0 {
		return entity.ErrBasFieldValue{
			Name: "quantity",
		}
	}

	return s.repo.UpdateProduct(ctx, id, req)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteProduct(ctx, id)
}
