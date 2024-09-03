package errs

import "errors"

// Authentication Errors
var (
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrPasswordIsEmpty             = errors.New("ErrPasswordIsEmpty")
	ErrPasswordIncorrect           = errors.New("ErrPasswordIncorrect")
	ErrUsernameIsEmpty             = errors.New("ErrUsernameIsEmpty")
	ErrEmailIsEmpty                = errors.New("ErrUsernameIsEmpty")
	ErrUsernameOrEmailIsEmpty      = errors.New("ErrUsernameOrEmailIsEmpty")
	ErrUsernameOrPasswordIsEmpty   = errors.New("ErrUsernameOrPasswordIsEmpty")
	ErrEmailOrPasswordIsEmpty      = errors.New("ErrEmailOrPasswordIsEmpty")
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
)

// Validation Errors
var (
	ErrInvalidData      = errors.New("ErrInvalidData")
	ErrValidationFailed = errors.New("ErrValidationFailed")
)

// Uniqueness Errors
var (
	ErrUsernameUniquenessFailed = errors.New("ErrUsernameUniquenessFailed")
	ErrEmailUniquenessFailed    = errors.New("ErrEmailUniquenessFailed")
)

// General Errors
var (
	ErrRecordNotFound     = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong = errors.New("ErrSomethingWentWrong")
)
