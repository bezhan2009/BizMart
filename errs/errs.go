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
	ErrPathParametrized = errors.New("ErrPathParametrized")
)

// Uniqueness Errors
var (
	ErrUsernameUniquenessFailed        = errors.New("ErrUsernameUniquenessFailed")
	ErrEmailUniquenessFailed           = errors.New("ErrEmailUniquenessFailed")
	ErrCategoryNameUniquenessFailed    = errors.New("ErrCategoryNameUniquenessFailed")
	ErrOrderStatusNameUniquenessFailed = errors.New("ErrOrderStatusNameUniquenessFailed")
)

// General Errors
var (
	ErrRecordNotFound      = errors.New("ErrRecordNotFound")
	ErrCategoryNotFound    = errors.New("ErrCategoryNotFound")
	ErrOrderStatusNotFound = errors.New("ErrOrderStatusNotFound")
	ErrSomethingWentWrong  = errors.New("ErrSomethingWentWrong")
)
