package repository

import (
	"BizMart/pkg/errs"
	"errors"
	"gorm.io/gorm"
)

func TranslateGormError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}

	return err
}
