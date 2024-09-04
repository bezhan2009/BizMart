package Order

import (
	"BizMart/errs"
	"BizMart/logger"
	"BizMart/models"
	"BizMart/pkg/controllers/handlers"
	"BizMart/pkg/repository/orderRepository"
	"BizMart/pkg/service/OrderService"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllOrderStatuses(c *gin.Context) {
	orderStatus, err := orderRepository.GetAllOrderStatuses()
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderStatus})
}

func GetOrderStatusByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handlers.HandleError(c, errs.ErrPathParametrized)
		return
	}

	orderStatus, err := orderRepository.GetOrderStatusByID(uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			handlers.HandleError(c, errs.ErrOrderStatusNotFound)
			return
		}

		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderStatus})
}

func GetOrderStatusByName(c *gin.Context) {
	name := c.Param("name")
	orderStatus, err := orderRepository.GetOrderStatusByName(name)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderStatus})
}

func CreateOrderStatus(c *gin.Context) {
	var orderStatus models.OrderStatus
	if err := c.ShouldBindJSON(&orderStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidData.Error()})
		return
	}

	orderStatusID, err := OrderService.CreateOrderStatus(orderStatus)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.CreateOrderStatus] successfully created new order status with ID %d", orderStatusID)

	c.JSON(http.StatusCreated, gin.H{"message": "order status created successfully"})
}

func UpdateOrderStatus(c *gin.Context) {
	orderStatusID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handlers.HandleError(c, errs.ErrValidationFailed)
		return
	}

	var OrdStat models.OrderStatus
	if err = c.BindJSON(&OrdStat); err != nil {
		handlers.HandleError(c, errs.ErrValidationFailed)
		return
	}

	OrdStat.ID = uint(orderStatusID)

	OrderStatusIDUpdated, err := OrderService.UpdateOrderStatus(uint(orderStatusID), OrdStat)
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateOrderStatus] successfully updated order status with ID %d", OrderStatusIDUpdated)

	c.JSON(http.StatusOK, gin.H{
		"message": "Order Status updated",
	})
}

func DeleteOrderStatus(c *gin.Context) {
	orderStatusID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.ErrPathParametrized,
		})
		return
	}

	orderStatus, err := orderRepository.GetOrderStatusByID(uint(orderStatusID))
	if err != nil {
		handlers.HandleError(c, err)
		return
	}

	if err := orderRepository.DeleteOrderStatus(orderStatus); err != nil {
		handlers.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order status deleted successfully"})
}
