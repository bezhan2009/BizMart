package repository

import (
	"BizMart/pkg/errs"
	"errors"
	"gorm.io/gorm"
	"strings"
)

func TranslateGormError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrRecordNotFound
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return errs.ErrDuplicateEntry
	}

	if errors.Is(err, gorm.ErrInvalidField) {
		return errs.ErrInvalidField
	}
	if errors.Is(err, gorm.ErrInvalidData) {
		return errs.ErrInvalidData
	}
	if errors.Is(err, gorm.ErrUnsupportedDriver) {
		return errs.ErrUnsupportedDriver
	}
	if errors.Is(err, gorm.ErrNotImplemented) {
		return errs.ErrNotImplemented
	}
	if isForeignKeyViolation(err) {
		return errs.ErrInvalidField
	}
	return err
}

func isForeignKeyViolation(err error) bool {
	return strings.Contains(err.Error(), "violates foreign key constraint")
}
