package errs

import (
	"errors"
	"github.com/jinzhu/gorm"
)

func TranslateGormError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecordNotFound
	}

	return err
}
