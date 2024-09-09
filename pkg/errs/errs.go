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
	ErrInvalidData       = errors.New("ErrInvalidData")
	ErrValidationFailed  = errors.New("ErrValidationFailed")
	ErrPathParametrized  = errors.New("ErrPathParametrized")
	ErrInvalidMinPrice   = errors.New("ErrInvalidMinPrice")
	ErrInvalidMaxPrice   = errors.New("ErrInvalidMaxPrice")
	ErrInvalidAmount     = errors.New("ErrInvalidAmount")
	ErrInsufficientFunds = errors.New("ErrInsufficientFunds")
	ErrInvalidCategory   = errors.New("ErrInvalidCategory")
	ErrInvalidStore      = errors.New("ErrInvalidStore")
	ErrInvalidID         = errors.New("ErrInvalidID")
	ErrInvalidProductID  = errors.New("ErrInvalidProductID")
	ErrInvalidStoreID    = errors.New("ErrInvalidProductID")
)

// Uniqueness Errors
var (
	ErrUsernameUniquenessFailed        = errors.New("ErrUsernameUniquenessFailed")
	ErrEmailUniquenessFailed           = errors.New("ErrEmailUniquenessFailed")
	ErrCategoryNameUniquenessFailed    = errors.New("ErrCategoryNameUniquenessFailed")
	ErrOrderStatusNameUniquenessFailed = errors.New("ErrOrderStatusNameUniquenessFailed")
	ErrStoreNameUniquenessFailed       = errors.New("ErrStoreNameUniquenessFailed")
)

// General Errors
var (
	ErrRecordNotFound      = errors.New("ErrRecordNotFound")
	ErrCategoryNotFound    = errors.New("ErrCategoryNotFound")
	ErrOrderStatusNotFound = errors.New("ErrOrderStatusNotFound")
	ErrSomethingWentWrong  = errors.New("ErrSomethingWentWrong")
	ErrProductNotFound     = errors.New("ErrProductNotFound")
	ErrStoreNotFound       = errors.New("ErrStoreNotFound")
	ErrUserNotFound        = errors.New("ErrUserNotFound")
	ErrDeleteFailed        = errors.New("ErrDeleteFailed")
)

// GORM Errors
var (
	ErrDuplicateEntry    = errors.New("ErrDuplicateEntry")
	ErrInvalidField      = errors.New("ErrInvalidField")
	ErrUnsupportedDriver = errors.New("ErrUnsupportedDriver")
	ErrNotImplemented    = errors.New("ErrNotImplemented")
)
