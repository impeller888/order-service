package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Firstname    string
	Lastname     string
	Fullname     string
	Age          *int
	IsMarried    *bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PasswordHash []byte
}

type UserRequest struct {
	Firstname    *string
	Lastname     *string
	Fullname     *string
	Age          *int
	IsMarried    *bool
	Password     *string
	PasswordHash []byte
}
