package errs

import "errors"

// GORM Errors
var (
	ErrDuplicateEntry    = errors.New("ErrDuplicateEntry")
	ErrInvalidField      = errors.New("ErrInvalidField")
	ErrUnsupportedDriver = errors.New("ErrUnsupportedDriver")
	ErrNotImplemented    = errors.New("ErrNotImplemented")
)
