package db

import "errors"

var (
	ErrDBError             = errors.New("db error")
	ErrMissedRequiredField = errors.New("missed reqired field")
)
