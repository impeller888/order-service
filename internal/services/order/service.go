package order

import (
	"context"
	"local/order-service/internal/entity"

	"github.com/google/uuid"
)

type OrderService struct {
	userSrv     userService
	orderRepo   orderRepository
	productRepo productRepository
}

func NewOrderService(
	userSrv userService,
	orderRepo orderRepository,
	productRepo productRepository,
) *OrderService {
	return &OrderService{
		userSrv:     userSrv,
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) GetOrderByID(ctx context.Context, id uuid.UUID) (*entity.Order, error) {
	return s.orderRepo.GetOrderByID(ctx, id)
}

func (s *OrderService) CreateOrder(ctx context.Context, req *entity.OrderRequest) (uuid.UUID, error) {
	if req.UserID == uuid.Nil {
		return uuid.Nil, &entity.ErrFieldRequired{
			Name: "user_id",
		}
	}
	// Проверяем что такой пользователь существует
	_, err := s.userSrv.GetUserByID(ctx, req.UserID)
	if err != nil {
		return uuid.Nil, entity.ErrUserNotExist
	}

	if len(req.Items) == 0 {
		return uuid.Nil, entity.ErrEmptyOrderItems
	}

	var totalCost uint

	// Проверяем позиции заказа
	for _, item := range req.Items {
		if item.ProductID == uuid.Nil {
			return uuid.Nil, &entity.ErrFieldRequired{
				Name: "product_id",
			}
		}
		if item.Price == nil {
			return uuid.Nil, &entity.ErrFieldRequired{
				Name: "price",
			}
		}
		// Проверяем наличие товара
		product, err := s.productRepo.GetProductByID(ctx, item.ProductID)
		if err != nil {
			return uuid.Nil, entity.ErrProductNotExist
		}
		// Проверяем доступность товара
		if *item.Quantity > product.Quantity {
			return uuid.Nil, entity.ErrNotEnoughProducts
		}
		diff := product.Quantity - *item.Quantity
		// Обновляем кол-во товара на складе
		err = s.productRepo.UpdateProduct(ctx, item.ProductID, &entity.ProductRequest{
			Quantity: &diff,
		})
		if err != nil {
			return uuid.Nil, err
		}
		totalCost += *item.Price
	}

	if req.TotalCost == nil {
		req.TotalCost = &totalCost
	}

	status := entity.OrderStatusPending
	req.Status = &status

	return s.orderRepo.CreateOrder(ctx, req)
}
