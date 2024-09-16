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

func GetAllUserOrders(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	orders, err := repository.GetAllOrderByUserID(userID)
	if err != nil {
		HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetOrderByID(c *gin.Context) {
	orderIdStr := c.Param("id")
	if orderIdStr == "" {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	orderID, err := strconv.ParseUint(orderIdStr, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	if orderID == 0 {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	order, err := repository.GetOrderByID(uint(orderID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrOrderNotFound)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func CreateOrder(c *gin.Context) {
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	var orderRequest models.OrderRequestJsonBind
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateOrder(orderRequest); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created successfully"})
}

func UpdateOrder(c *gin.Context) {
	orderIdStr := c.Param("id")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		HandleError(c, errs.ErrUnauthorized)
		return
	}

	var orderRequest models.OrderRequestJsonBind
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.UpdateOrder(uint(orderId), orderRequest); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created successfully"})
}

func DeleteOrder(c *gin.Context) {
	orderIdStr := c.Param("id")
	if orderIdStr == "" {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	orderID, err := strconv.ParseUint(orderIdStr, 10, 64)
	if err != nil {
		HandleError(c, errs.ErrInvalidOrderID)
		return
	}

	if err := repository.DeleteOrder(uint(orderID)); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}
