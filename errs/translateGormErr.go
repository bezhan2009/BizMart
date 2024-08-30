package errs

import (
	"errors"
	"gorm.io/gorm"
)

func TranslateGormError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecordNotFound
	}

	return err
}
