package user

import (
	"context"
	"local/order-service/internal/entity"

	"github.com/google/uuid"
)

type userService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	RegisterUser(ctx context.Context, req *entity.UserRequest) (uuid.UUID, error)
	UpdateUser(ctx context.Context, id uuid.UUID, req *entity.UserRequest) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
