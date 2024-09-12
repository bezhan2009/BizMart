package errs

import "errors"

// Uniqueness Errors
var (
	ErrUsernameUniquenessFailed        = errors.New("ErrUsernameUniquenessFailed")
	ErrAccountNumberUniquenessFailed   = errors.New("ErrAccountNumberUniquenessFailed")
	ErrAddressNameUniquenessFailed     = errors.New("ErrAddressNameUniquenessFailed")
	ErrEmailUniquenessFailed           = errors.New("ErrEmailUniquenessFailed")
	ErrCategoryNameUniquenessFailed    = errors.New("ErrCategoryNameUniquenessFailed")
	ErrOrderStatusNameUniquenessFailed = errors.New("ErrOrderStatusNameUniquenessFailed")
	ErrStoreNameUniquenessFailed       = errors.New("ErrStoreNameUniquenessFailed")
)
