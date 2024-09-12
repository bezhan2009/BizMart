package errs

import "errors"

// Authentication Errors
var (
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrPasswordIsEmpty             = errors.New("ErrPasswordIsEmpty")
	ErrPasswordIncorrect           = errors.New("ErrPasswordIncorrect")
	ErrUsernameIsEmpty             = errors.New("ErrUsernameIsEmpty")
	ErrEmailIsEmpty                = errors.New("ErrEmailIsEmpty")
	ErrUsernameOrEmailIsEmpty      = errors.New("ErrUsernameOrEmailIsEmpty")
	ErrUsernameOrPasswordIsEmpty   = errors.New("ErrUsernameOrPasswordIsEmpty")
	ErrEmailOrPasswordIsEmpty      = errors.New("ErrEmailOrPasswordIsEmpty")
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
	ErrUnauthorized                = errors.New("ErrUnauthorized")
)

// Validation Errors
var (
	ErrInvalidData          = errors.New("ErrInvalidData")
	ErrValidationFailed     = errors.New("ErrValidationFailed")
	ErrPathParametrized     = errors.New("ErrPathParametrized")
	ErrInvalidMinPrice      = errors.New("ErrInvalidMinPrice")
	ErrInvalidMaxPrice      = errors.New("ErrInvalidMaxPrice")
	ErrInvalidAmount        = errors.New("ErrInvalidAmount")
	ErrInvalidPrice         = errors.New("ErrInvalidPrice")
	ErrInsufficientFunds    = errors.New("ErrInsufficientFunds")
	ErrInvalidCategory      = errors.New("ErrInvalidCategory")
	ErrInvalidStore         = errors.New("ErrInvalidStore")
	ErrInvalidID            = errors.New("ErrInvalidID")
	ErrInvalidAddressID     = errors.New("ErrInvalidAddressID")
	ErrInvalidProductID     = errors.New("ErrInvalidProductID")
	ErrInvalidStoreID       = errors.New("ErrInvalidStoreID")
	ErrInvalidStoreReviewID = errors.New("ErrInvalidStoreReviewID")
	ErrInvalidComment       = errors.New("ErrInvalidComment")
	ErrInvalidRating        = errors.New("ErrInvalidRating")
	ErrInvalidTitle         = errors.New("ErrInvalidTitle")
	ErrInvalidAddressName   = errors.New("ErrInvalidAddressName")
	ErrInvalidDescription   = errors.New("ErrInvalidDescription")
)

// Uniqueness Errors
var (
	ErrUsernameUniquenessFailed        = errors.New("ErrUsernameUniquenessFailed")
	ErrAddressNameUniquenessFailed     = errors.New("ErrAddressNameUniquenessFailed")
	ErrEmailUniquenessFailed           = errors.New("ErrEmailUniquenessFailed")
	ErrCategoryNameUniquenessFailed    = errors.New("ErrCategoryNameUniquenessFailed")
	ErrOrderStatusNameUniquenessFailed = errors.New("ErrOrderStatusNameUniquenessFailed")
	ErrStoreNameUniquenessFailed       = errors.New("ErrStoreNameUniquenessFailed")
)

// General Errors
var (
	ErrAddressNotFound     = errors.New("ErrAddressNotFound")
	ErrRecordNotFound      = errors.New("ErrRecordNotFound")
	ErrCategoryNotFound    = errors.New("ErrCategoryNotFound")
	ErrOrderStatusNotFound = errors.New("ErrOrderStatusNotFound")
	ErrSomethingWentWrong  = errors.New("ErrSomethingWentWrong")
	ErrProductNotFound     = errors.New("ErrProductNotFound")
	ErrStoreNotFound       = errors.New("ErrStoreNotFound")
	ErrUserNotFound        = errors.New("ErrUserNotFound")
	ErrDeleteFailed        = errors.New("ErrDeleteFailed")
	ErrFetchingProducts    = errors.New("ErrFetchingProducts")
	ErrNoProductsFound     = errors.New("ErrNoProductsFound")
	ErrStoreReviewNotFound = errors.New("ErrStoreReviewNotFound")
)

// GORM Errors
var (
	ErrDuplicateEntry    = errors.New("ErrDuplicateEntry")
	ErrInvalidField      = errors.New("ErrInvalidField")
	ErrUnsupportedDriver = errors.New("ErrUnsupportedDriver")
	ErrNotImplemented    = errors.New("ErrNotImplemented")
)
