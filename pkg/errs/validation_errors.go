package errs

import "errors"

// Validation Errors
var (
	ErrInvalidData              = errors.New("ErrInvalidData")
	ErrValidationFailed         = errors.New("ErrValidationFailed")
	ErrPathParametrized         = errors.New("ErrPathParametrized")
	ErrInvalidMinPrice          = errors.New("ErrInvalidMinPrice")
	ErrInvalidMaxPrice          = errors.New("ErrInvalidMaxPrice")
	ErrInvalidAmount            = errors.New("ErrInvalidAmount")
	ErrInvalidPrice             = errors.New("ErrInvalidPrice")
	ErrInsufficientFunds        = errors.New("ErrInsufficientFunds")
	ErrInvalidCategory          = errors.New("ErrInvalidCategory")
	ErrInvalidStore             = errors.New("ErrInvalidStore")
	ErrInvalidID                = errors.New("ErrInvalidID")
	ErrInvalidFeaturedProductID = errors.New("ErrInvalidFeaturedProductID")
	ErrInvalidAccountID         = errors.New("ErrInvalidAccountID")
	ErrInvalidAddressID         = errors.New("ErrInvalidAddressID")
	ErrInvalidProductID         = errors.New("ErrInvalidProductID")
	ErrInvalidProductReviewID   = errors.New("ErrInvalidProductReviewID")
	ErrInvalidStoreID           = errors.New("ErrInvalidStoreID")
	ErrInvalidStoreReviewID     = errors.New("ErrInvalidStoreReviewID")
	ErrInvalidComment           = errors.New("ErrInvalidComment")
	ErrInvalidContent           = errors.New("ErrInvalidContent")
	ErrInvalidRating            = errors.New("ErrInvalidRating")
	ErrInvalidTitle             = errors.New("ErrInvalidTitle")
	ErrInvalidAddressName       = errors.New("ErrInvalidAddressName")
	ErrInvalidAccountNumber     = errors.New("ErrInvalidAccountNumber")
	ErrInvalidDescription       = errors.New("ErrInvalidDescription")
)
