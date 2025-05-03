package user

import (
	"local/order-service/internal/entity"
	"time"

	"github.com/google/uuid"
)

type user struct {
	ID           uuid.UUID
	Firstname    string
	Lastname     string
	Fullname     string
	Age          *int
	IsMarried    *bool
	PasswordHash []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *user) fillFromEntity(entity *entity.User) {
	u.ID = entity.ID
	u.Firstname = entity.Firstname
	u.Lastname = entity.Lastname
	u.Fullname = entity.Fullname
	u.Age = entity.Age
	u.IsMarried = entity.IsMarried
	u.PasswordHash = entity.PasswordHash
	u.CreatedAt = entity.CreatedAt
	u.UpdatedAt = entity.UpdatedAt
}

func (u *user) toEntity() *entity.User {
	user := &entity.User{
		ID:           u.ID,
		Firstname:    u.Firstname,
		Lastname:     u.Lastname,
		Fullname:     u.Fullname,
		Age:          u.Age,
		IsMarried:    u.IsMarried,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}

	return user
}
