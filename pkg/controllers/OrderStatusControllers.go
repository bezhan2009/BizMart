package controllers

import (
	"BizMart/errs"
	"BizMart/logger"
	"BizMart/models"
	"BizMart/pkg/repository/order"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllOrderStatusses(c *gin.Context) {
	orderStatus, err := order.GetAllOrderStatuses()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderStatus})
}

func GetOrderStatusByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(c, err)
		return
	}

	orderStatus, err := order.GetOrderStatusByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderStatus})
}

func GetOrderStatusByName(c *gin.Context) {
	name := c.Param("name")
	orderStatus, err := order.GetOrderStatusByName(name)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderStatus})
}

func CreateOrderStatus(c *gin.Context) {
	var orderStatus models.OrderStatus
	if err := c.ShouldBindJSON(&orderStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidData})
		return
	}

	orderStatusID, err := order.CreateOrderStatus(orderStatus)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.CreateOrderStatus] successfully created new order status with ID %d", orderStatusID)

	c.JSON(http.StatusCreated, gin.H{"message": "order status created successfully"})
}

func UpdateOrderStatus(c *gin.Context) {
	orderStatusID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	var OrdStat models.OrderStatus
	if err = c.BindJSON(&OrdStat); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	OrdStat.ID = uint(orderStatusID)

	OrderStatusIDUpdated, err := order.UpdateOrderStatus(OrdStat)
	if err != nil {
		handleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateOrderStatus] successfully updated order status with ID %d", OrderStatusIDUpdated)

	c.JSON(http.StatusOK, gin.H{
		"message": "Order Status updated",
	})
}
