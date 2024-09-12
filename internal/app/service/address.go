package service

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/errs"
	"github.com/gin-gonic/gin"
)

func ValidateAddress(HandleError func(ctx *gin.Context, err error), addressData models.Address, c *gin.Context) error {
	if len(addressData.AddressName) <= 2 {
		HandleError(c, errs.ErrInvalidAddressName)
		return errs.ErrInvalidAddressName
	}

	if len(addressData.AddressName) > 50 {
		HandleError(c, errs.ErrInvalidAddressName)
		return errs.ErrInvalidAddressName
	}

	if addressData.IsDeleted {
		HandleError(c, errs.ErrPermissionDenied)
		return errs.ErrPermissionDenied
	}

	return nil
}
