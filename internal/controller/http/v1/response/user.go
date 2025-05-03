package response

import (
	"local/order-service/internal/entity"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Fullname  string    `json:"fullname"`
	Age       *int      `json:"age"`
	IsMarried *bool     `json:"is_married"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) FillFromEntity(entity *entity.User) {
	u.ID = entity.ID
	u.Firstname = entity.Firstname
	u.Lastname = entity.Lastname
	u.Fullname = entity.Fullname
	u.Age = entity.Age
	u.IsMarried = entity.IsMarried
	u.CreatedAt = entity.CreatedAt
	u.UpdatedAt = entity.UpdatedAt
}
