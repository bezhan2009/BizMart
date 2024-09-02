package controllers

import (
	"BizMart/errs"
	"BizMart/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(c *gin.Context, err error) {
	if errors.Is(err, errs.ErrInvalidData) {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidDataCustom})
	} else if errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrPermissionDenied) {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		logger.Error.Printf("Err: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errs.ErrSomethingWentWrong.Error()})
	}
}
