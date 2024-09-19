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

// GetUserPayments godoc
// @Summary Get user payments
// @Description Get all payments of the authenticated user
// @Tags Payments
// @Accept  json
// @Produce  json
// @Success 200 {object} models.DefaultResponse "payments"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /payments [get]
// @Security ApiKeyAuth
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

// GetPaymentByID godoc
// @Summary Get payment by ID
// @Description Get a payment by its ID for the authenticated user
// @Tags Payments
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Success 200 {object} models.Payment "payment"
// @Failure 400 {object} models.ErrorResponse "Invalid Payment ID"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Payment Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /payments/{id} [get]
// @Security ApiKeyAuth
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

// CreatePayment godoc
// @Summary Create a new payment
// @Description Create a new payment for the authenticated user
// @Tags Payments
// @Accept  json
// @Produce  json
// @Param payment body models.Payment true "Payment Data"
// @Success 201 {object} models.DefaultResponse "Payment Created Successfully"
// @Failure 400 {object} models.ErrorResponse "Validation Failed"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /payments [post]
// @Security ApiKeyAuth
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

// UpdatePayment godoc
// @Summary Update payment by ID
// @Description Update an existing payment by its ID for the authenticated user
// @Tags Payments
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Param payment body models.PaymentRequest true "Payment Data"
// @Success 200 {object} models.DefaultResponse "Payment Updated Successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid Payment ID or Validation Failed"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Payment Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /payments/{id} [put]
// @Security ApiKeyAuth
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

// DeletePayment godoc
// @Summary Delete payment by ID
// @Description Delete a payment by its ID for the authenticated user
// @Tags Payments
// @Accept  json
// @Produce  json
// @Param id path string true "Payment ID"
// @Success 200 {object} models.DefaultResponse "Payment Deleted Successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid Payment ID"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 404 {object} models.ErrorResponse "Payment Not Found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /payments/{id} [delete]
// @Security ApiKeyAuth
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
