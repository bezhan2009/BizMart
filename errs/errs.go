package errs

import "errors"

var (
	ErrUsernameUniquenessFailed    = errors.New("ErrUsernameUniquenessFailed")
	ErrIncorrectUsernameOrPassword = errors.New("ErrIncorrectUsernameOrPassword")
	ErrEmailUniquenessFailed       = errors.New("ErrEmailUniquenessFailed")
	ErrRecordNotFound              = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong          = errors.New("ErrSomethingWentWrong")
	ErrInvalidData                 = errors.New("ErrInvalidData")
	ErrInvalidDataCustom           = errors.New("ErrInvalidDataCustom")
	ErrPasswordIsEmpty             = errors.New("ErrPasswordIsEmpty")
	ErrPasswordIncorrect           = errors.New("ErrPasswordIncorrect")
	ErrEmailIsEmpty                = errors.New("ErrEmailIsEmpty")
	ErrUsernameIsEmpty             = errors.New("ErrUsernameIsEmpty")
	ErrUsernameOrEmailIsEmpty      = errors.New("ErrUsernameOrEmailIsEmpty")
	ErrUsernameOrPasswordIsEmpty   = errors.New("ErrUsernameOrPasswordIsEmpty")
	ErrEmailOrPasswordIsEmpty      = errors.New("ErrEmailOrPasswordIsEmpty")
	ErrPermissionDenied            = errors.New("ErrPermissionDenied")
)
