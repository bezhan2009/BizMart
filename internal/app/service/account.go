package service

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/errs"
	"github.com/gin-gonic/gin"
)

func ValidateAccount(HandleError func(ctx *gin.Context, err error), accountData models.Account, c *gin.Context) error {
	if len(accountData.AccountNumber) < 4 {
		HandleError(c, errs.ErrInvalidAccountNumber)
		return errs.ErrInvalidAccountNumber
	}

	if len(accountData.AccountNumber) > 50 {
		HandleError(c, errs.ErrInvalidAccountNumber)
		return errs.ErrInvalidAccountNumber
	}

	if accountData.IsDeleted {
		HandleError(c, errs.ErrPermissionDenied)
		return errs.ErrPermissionDenied
	}

	return nil
}
