package repository

import "errors"

var (
	ErrMissingDBClient = errors.New("db client should not be nil")
)
