package user

import (
	"context"
	"local/order-service/internal/entity"

	"github.com/google/uuid"
)

type userRepository interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	CreateUser(ctx context.Context, req *entity.UserRequest) (uuid.UUID, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, req *entity.UserRequest) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
