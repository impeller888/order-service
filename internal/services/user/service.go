package user

import (
	"context"
	"local/order-service/internal/entity"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type UserService struct {
	repo userRepository
}

func NewUserService(repo userRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, req *entity.UserRequest) (uuid.UUID, error) {
	var err error

	if req.Firstname == nil || *req.Firstname == "" {
		return uuid.Nil, &entity.ErrFieldRequired{
			Name: "firstname",
		}
	}
	if req.Lastname == nil || *req.Lastname == "" {
		return uuid.Nil, &entity.ErrFieldRequired{
			Name: "lastname",
		}
	}
	if req.IsMarried == nil {
		return uuid.Nil, &entity.ErrFieldRequired{
			Name: "is_married",
		}
	}
	if req.Fullname == nil || *req.Fullname == "" {
		fullname := *req.Firstname + " " + *req.Lastname
		req.Fullname = &fullname
	}

	if req.Age == nil {
		return uuid.Nil, &entity.ErrFieldRequired{
			Name: "age",
		}
	} else if *req.Age < 18 {
		return uuid.Nil, entity.ErrInsufficientAge
	}

	if req.Password == nil {
		return uuid.Nil, &entity.ErrFieldRequired{
			Name: "password",
		}
	}

	if utf8.RuneCountInString(*req.Password) < 8 {
		return uuid.Nil, entity.ErrPasswordTooShort
	}

	req.PasswordHash, err = bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, entity.ErrBadPasswordFormat
	}

	return s.repo.CreateUser(ctx, req)
}

func (s *UserService) GetUserByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}

func (s *UserService) UpdateUser(ctx context.Context, userID uuid.UUID, req *entity.UserRequest) error {
	return s.repo.UpdateUser(ctx, userID, req)
}

func (s *UserService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return s.repo.DeleteUser(ctx, userID)
}
