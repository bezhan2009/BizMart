package errs

import "errors"

// Authentication Errors
var (
	ErrInvalidCredentials          = errors.New("invalid credentials")
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
	ErrEmailIsRequired             = errors.New("email is required")
	ErrUsernameIsRequired          = errors.New("username is required")
	ErrFirstNameIsRequired         = errors.New("first name is required")
	ErrLastNameIsRequired          = errors.New("last name is required")
	ErrPasswordIsRequired          = errors.New("password is required")
	ErrAppLoginIsRequired          = errors.New("app login is required")
)
