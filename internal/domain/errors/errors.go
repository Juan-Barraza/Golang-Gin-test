package errors

import "errors"

var (
	ErrInvalidType = errors.New("invalid type")
	ErrInvalidDate = errors.New("invalid date")
)
