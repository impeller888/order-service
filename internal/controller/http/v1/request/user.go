package request

import "local/order-service/internal/entity"

type UserRequest struct {
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Fullname  *string `json:"fullname"`
	Age       *int    `json:"age"`
	IsMarried *bool   `json:"is_married"`
	Password  *string `json:"password"`
}

func (u *UserRequest) ToEntity() *entity.UserRequest {
	return &entity.UserRequest{
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Fullname:  u.Fullname,
		Age:       u.Age,
		IsMarried: u.IsMarried,
		Password:  u.Password,
	}
}
