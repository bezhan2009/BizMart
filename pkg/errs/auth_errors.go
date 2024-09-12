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
