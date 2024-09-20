package service

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/errs"
	"github.com/gin-gonic/gin"
)

func ValidateProduct(HandleError func(ctx *gin.Context, err error), productData models.Product, c *gin.Context, isUpdate bool) error {
	if productData.Amount <= 0 {
		HandleError(c, errs.ErrInvalidAmount)
		return errs.ErrInvalidAmount
	}

	if productData.Price <= 0 {
		HandleError(c, errs.ErrInvalidPrice)
		return errs.ErrInvalidPrice
	}

	if productData.CategoryID <= 0 {
		HandleError(c, errs.ErrInvalidCategory)
		return errs.ErrInvalidCategory
	}

	if len(productData.Title) <= 5 {
		HandleError(c, errs.ErrInvalidTitle)
		return errs.ErrInvalidTitle
	}

	if len(productData.Description) <= 5 {
		HandleError(c, errs.ErrInvalidDescription)
		return errs.ErrInvalidDescription
	}

	if productData.Views > 0 && !isUpdate {
		HandleError(c, errs.ErrPermissionDenied)
		return errs.ErrPermissionDenied
	}

	return nil
}
