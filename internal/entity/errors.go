package entity

import (
	"errors"
)

type Error error

var (
	// Main
	ErrBadIDFormat Error = errors.New("bad id format")
	ErrNotFound    Error = errors.New("not found")

	// User
	ErrUserNotExist      Error = errors.New("user not exist")
	ErrPasswordTooShort  Error = errors.New("password too short")
	ErrBadPasswordFormat Error = errors.New("bad password format")
	ErrInsufficientAge   Error = errors.New("insufficient age")
	// Order
	ErrEmptyOrderItems   Error = errors.New("empty order items")
	ErrNotEnoughProducts Error = errors.New("not enough products")
	ErrProductNotExist   Error = errors.New("product not exist")
)

type ErrBasFieldValue struct {
	Name string
}

type ErrFieldRequired struct {
	Name string
}

func (e ErrFieldRequired) Error() string {
	return e.Name + " is required"
}

func (e ErrBasFieldValue) Error() string {
	return "bad value of " + e.Name
}
