package errs

import "errors"

// General Errors
var (
	ErrAddressNotFound         = errors.New("ErrAddressNotFound")
	ErrAccountNotFound         = errors.New("ErrAccountNotFound")
	ErrFeaturedProductNotFound = errors.New("ErrFeaturedProductNotFound")
	ErrRecordNotFound          = errors.New("ErrRecordNotFound")
	ErrCategoryNotFound        = errors.New("ErrCategoryNotFound")
	ErrOrderStatusNotFound     = errors.New("ErrOrderStatusNotFound")
	ErrSomethingWentWrong      = errors.New("ErrSomethingWentWrong")
	ErrProductNotFound         = errors.New("ErrProductNotFound")
	ErrStoreNotFound           = errors.New("ErrStoreNotFound")
	ErrUserNotFound            = errors.New("ErrUserNotFound")
	ErrDeleteFailed            = errors.New("ErrDeleteFailed")
	ErrFetchingProducts        = errors.New("ErrFetchingProducts")
	ErrNoProductsFound         = errors.New("ErrNoProductsFound")
	ErrStoreReviewNotFound     = errors.New("ErrStoreReviewNotFound")
)
