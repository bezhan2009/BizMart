package controllers

import (
	"BizMart/pkg/errs"
	"BizMart/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	if errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrCategoryNameUniquenessFailed) ||
		errors.Is(err, errs.ErrOrderStatusNameUniquenessFailed) ||
		errors.Is(err, errs.ErrPathParametrized) ||
		errors.Is(err, errs.ErrInvalidProductID) ||
		errors.Is(err, errs.ErrInvalidMinPrice) ||
		errors.Is(err, errs.ErrInvalidMaxPrice) ||
		errors.Is(err, errs.ErrInvalidCategory) ||
		errors.Is(err, errs.ErrEmailIsEmpty) ||
		errors.Is(err, errs.ErrPasswordIsEmpty) ||
		errors.Is(err, errs.ErrUsernameIsEmpty) ||
		errors.Is(err, errs.ErrInvalidStore) ||
		errors.Is(err, errs.ErrValidationFailed) ||
		errors.Is(err, errs.ErrStoreNameUniquenessFailed) ||
		errors.Is(err, errs.ErrDeleteFailed) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrPermissionDenied) {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrRecordNotFound) ||
		errors.Is(err, errs.ErrCategoryNotFound) ||
		errors.Is(err, errs.ErrOrderStatusNotFound) ||
		errors.Is(err, errs.ErrProductNotFound) ||
		errors.Is(err, errs.ErrStoreNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		logger.Error.Printf("Err: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
	}
}
