package order

import (
	"context"
	"fmt"
	"local/order-service/internal/entity"
	"local/order-service/internal/repo/db"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*entity.Order, error) {
	raw := &order{}
	err := r.db.QueryRow(ctx, "SELECT id, user_id, total_cost, status, created_at, updated_at FROM orders WHERE id = $1", orderID).Scan(
		&raw.ID, &raw.UserID, &raw.TotalCost, &raw.Status, &raw.CreatedAt, &raw.UpdatedAt)
	if err != nil {
		return nil, err
	}

	order := raw.toEntity()

	items, err := r.getOrderItems(ctx, orderID)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return order, nil

}

func (r *OrderRepository) CreateOrder(ctx context.Context, req *entity.OrderRequest) (uuid.UUID, error) {

	var id uuid.UUID

	if req.UserID == uuid.Nil {
		return uuid.Nil, db.ErrMissedRequiredField
	}

	if req.Status == nil {
		return uuid.Nil, db.ErrMissedRequiredField
	}

	var (
		namedArgs = pgx.NamedArgs{
			"user_id": req.UserID,
			"status":  *req.Status,
		}
		colNames, argNames []string
	)

	if req.TotalCost != nil {
		namedArgs["total_cost"] = *req.TotalCost
	} else {
		namedArgs["total_cost"] = 0
	}

	for col, _ := range namedArgs {
		colNames = append(colNames, col)
		argNames = append(argNames, "@"+col)
	}

	err := r.db.QueryRow(ctx, fmt.Sprintf("INSERT INTO orders (%s) VALUES (%s) RETURNING id", strings.Join(colNames, ","), strings.Join(argNames, ",")), namedArgs).
		Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	err = r.createOrderItems(ctx, id, req.Items)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *OrderRepository) getOrderItems(ctx context.Context, orderID uuid.UUID) ([]*entity.OrderItem, error) {
	rows, err := r.db.Query(ctx, "SELECT order_id, product_id, description, price, quantity, created_at, updated_at FROM order_items WHERE order_id = $1", orderID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*entity.OrderItem
	for rows.Next() {
		raw := &orderItem{}
		err = rows.Scan(&raw.OrderID, &raw.ProductID, &raw.Description, &raw.Price, &raw.Quantity, &raw.CreatedAt, &raw.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, raw.toEntity())
	}

	return items, nil
}

func (r *OrderRepository) createOrderItems(ctx context.Context, orderID uuid.UUID, items []*entity.OrderItemRequest) error {
	for _, item := range items {
		if item.ProductID == uuid.Nil || item.Description == nil || item.Price == nil || item.Quantity == nil {
			return db.ErrMissedRequiredField
		}
		_, err := r.db.Exec(ctx, `INSERT INTO order_items (order_id, product_id, description, quantity, price) VALUES ($1, $2, $3, $4, $5)`,
			orderID, item.ProductID, item.Description, item.Quantity, item.Price)
		if err != nil {
			return err
		}
	}

	return nil
}

// func (o *OrderRepository) AddOrderItem(ctx context.Context, item *entity.OrderItem) error {
// 	raw := &orderItem{}
// 	raw.fillFromEntity(item)
// 	_, err := r.db.Exec(ctx, "INSERT INTO order_items (order_id, product_id, quantity, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
// 		raw.OrderID, raw.ProductID, raw.Quantity, raw.Price, time.Now(), time.Now())
// 	return err
// }

// func (o *OrderRepository) RemoveOrderItem(ctx context.Context, orderID, productID uuid.UUID) error {
// 	_, err := r.pool.Exec(ctx, "DELETE FROM order_items WHERE order_i = $1 AND product_id = $2", orderID, productID)
// 	return err
// }
