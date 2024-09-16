package errs

import "errors"

// General Errors
var (
	ErrAddressNotFound         = errors.New("ErrAddressNotFound")
	ErrProductReviewNotFound   = errors.New("ErrProductReviewNotFound")
	ErrAccountNotFound         = errors.New("ErrAccountNotFound")
	ErrFeaturedProductNotFound = errors.New("ErrFeaturedProductNotFound")
	ErrRecordNotFound          = errors.New("ErrRecordNotFound")
	ErrOrderNotFound           = errors.New("ErrOrderNotFound")
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
