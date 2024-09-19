package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/internal/controllers/middlewares"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserPayments(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	payments, err := repository.GetAllUserPayments(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"payments": payments})
}

func GetPaymentByID(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	paymentStrID := c.Param("id")
	if paymentStrID == "" {
		HandleError(c, errs.ErrInvalidPaymentID)
		return
	}

	paymentID, err := strconv.ParseUint(paymentStrID, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidPaymentID)
		return
	}

	payment, err := repository.GetPaymentByID(uint(paymentID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrPaymentNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment": payment})
}

func CreatePayment(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	var payment models.Payment
	if err := c.ShouldBind(&payment); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	payment.UserID = userID

	if err := service.ValidatePayment(HandleError, &payment, c); err != nil {
		return
	}

	if err := service.CreatePayment(payment); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "payment created successfully"})
}

func UpdatePayment(c *gin.Context) {
	paymentStrID := c.Param("id")
	if paymentStrID == "" {
		HandleError(c, errs.ErrInvalidPaymentID)
		return
	}

	paymentID, err := strconv.ParseUint(paymentStrID, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidPaymentID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	paymentData, err := repository.GetPaymentByID(uint(paymentID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrPaymentNotFound)
			return
		}
		HandleError(c, err)
		return
	}

	if userID != paymentData.UserID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	var payment models.Payment
	if err := c.ShouldBind(&payment); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	payment.UserID = userID
	payment.ID = uint(paymentID)

	if err = service.ValidatePayment(HandleError, &payment, c); err != nil {
		return
	}

	if err = repository.UpdatePayment(payment); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrPaymentNotFound)
			return
		}
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment updated successfully"})
}

func DeletePayment(c *gin.Context) {
	paymentStrID := c.Param("id")
	if paymentStrID == "" {
		HandleError(c, errs.ErrInvalidPaymentID)
		return
	}

	paymentID, err := strconv.ParseUint(paymentStrID, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidPaymentID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	paymentData, err := repository.GetPaymentByID(uint(paymentID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrPaymentNotFound)
			return
		}
		HandleError(c, err)
		return
	}

	if userID != paymentData.UserID {
		HandleError(c, errs.ErrPermissionDenied)
		return
	}

	if err = repository.DeletePayment(paymentData); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment deleted successfully"})
}
