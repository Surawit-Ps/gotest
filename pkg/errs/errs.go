package errs

import "errors"

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrDatabase    = errors.New("database error")
	ErrStatusUnchanged = errors.New("status is already in the desired state")
)