package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"BizMart/pkg/logger"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllOrderStatuses godoc
// @Summary Get all order statuses
// @Description Fetches all order statuses from the database.
// @Tags order status
// @Accept  json
// @Produce  json
// @Success 200 {object} models.OrderStatus "Returns a list of order statuses"
// @Failure 500 {object} models.ErrorResponse
// @Router /order/status [get]
func GetAllOrderStatuses(c *gin.Context) {
	// Получаем все статусы заказов
	orderStatus, err := repository.GetAllOrderStatuses()
	if err != nil {
		HandleError(c, err)
		return
	}

	fmt.Println(orderStatus)

	// Возвращаем результат
	c.JSON(http.StatusOK, gin.H{"data": orderStatus})
}

// GetOrderStatusByID godoc
// @Summary Get order status by ID
// @Description Fetches a specific order status by its ID.
// @Tags order status
// @Accept  json
// @Produce  json
// @Param id path int true "Order Status ID"
// @Success 200 {object} models.OrderStatus "Returns the order status"
// @Failure 400 {object} models.ErrorResponse "Invalid ID"
// @Failure 404 {object} models.ErrorResponse "Order status not found"
// @Failure 500 {object} models.ErrorResponse
// @Router /order/status/{id} [get]
func GetOrderStatusByID(c *gin.Context) {
	// Преобразование параметра ID в число
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID) // Точная ошибка
		return
	}

	// Получение статуса по ID
	orderStatus, err := repository.GetOrderStatusByID(uint(id))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrOrderStatusNotFound)
			return
		}
		HandleError(c, err)
		return
	}

	// Возвращаем результат
	c.JSON(http.StatusOK, gin.H{"data": orderStatus})
}

// GetOrderStatusByName godoc
// @Summary Get order status by name
// @Description Fetches a specific order status by its name.
// @Tags order status
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param name path string true "Order Status Name"
// @Success 200 {object} models.OrderStatus "Returns the order status"
// @Failure 404 {object} models.ErrorResponse "Order status not found"
// @Failure 500 {object} models.ErrorResponse
// @Router /order/status/name/{name} [get]
func GetOrderStatusByName(c *gin.Context) {
	// Получаем имя из параметров
	name := c.Param("name")

	// Поиск статуса по имени
	orderStatus, err := repository.GetOrderStatusByName(name)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrOrderStatusNotFound)
			return
		}
		HandleError(c, err)
		return
	}

	// Возвращаем результат
	c.JSON(http.StatusOK, gin.H{"data": orderStatus})
}

// CreateOrderStatus godoc
// @Summary Create a new order status
// @Description Creates a new order status.
// @Tags order status
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param order_status body models.OrderStatusRequest true "Order Status data"
// @Success 201 {object} models.DefaultResponse "Order status created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 500 {object} models.ErrorResponse
// @Router /order/status [post]
func CreateOrderStatus(c *gin.Context) {
	var orderStatus models.OrderStatus
	// Проверяем корректность входных данных
	if err := c.ShouldBindJSON(&orderStatus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrInvalidData.Error()})
		return
	}

	// Создаем новый статус заказа
	orderStatusID, err := service.CreateOrderStatus(orderStatus)
	if err != nil {
		HandleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.CreateOrderStatus] successfully created new order status with ID %d", orderStatusID)

	// Возвращаем успешный результат
	c.JSON(http.StatusCreated, gin.H{"message": "Order status created successfully"})
}

// UpdateOrderStatus godoc
// @Summary Update an existing order status
// @Description Updates the details of an existing order status.
// @Tags order status
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Order Status ID"
// @Param order_status body models.OrderStatusRequest true "Updated order status data"
// @Success 200 {object} models.DefaultResponse "Order status updated successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input"
// @Failure 404 {object} models.ErrorResponse "Order status not found"
// @Failure 500 {object} models.ErrorResponse
// @Router /order/status/{id} [put]
func UpdateOrderStatus(c *gin.Context) {
	// Преобразование параметра ID в число
	orderStatusID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		HandleError(c, errs.ErrInvalidID) // Точная ошибка
		return
	}

	var OrdStat models.OrderStatus
	// Проверяем корректность входных данных
	if err = c.BindJSON(&OrdStat); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	OrdStat.ID = uint(orderStatusID)

	// Обновляем статус заказа
	OrderStatusIDUpdated, err := service.UpdateOrderStatus(uint(orderStatusID), OrdStat)
	if err != nil {
		HandleError(c, err)
		return
	}

	logger.Info.Printf("[controllers.UpdateOrderStatus] successfully updated order status with ID %d", OrderStatusIDUpdated)

	// Возвращаем успешный результат
	c.JSON(http.StatusOK, gin.H{
		"message": "Order Status updated",
	})
}

// DeleteOrderStatus godoc
// @Summary Delete an order status
// @Description Deletes a specific order status by its ID.
// @Tags order status
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id path int true "Order Status ID"
// @Success 200 {object} models.DefaultResponse "Order status deleted successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid ID"
// @Failure 404 {object} models.ErrorResponse "Order status not found"
// @Failure 500 {object} models.ErrorResponse
// @Router /order/status/{id} [delete]
func DeleteOrderStatus(c *gin.Context) {
	// Преобразование параметра ID в число
	orderStatusID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errs.ErrInvalidID, // Точная ошибка
		})
		return
	}

	// Проверяем наличие статуса заказа
	orderStatus, err := repository.GetOrderStatusByID(uint(orderStatusID))
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrOrderStatusNotFound)
			return
		}
		HandleError(c, err)
		return
	}

	// Удаляем статус заказа
	if err := repository.DeleteOrderStatus(orderStatus); err != nil {
		HandleError(c, err)
		return
	}

	// Возвращаем успешный результат
	c.JSON(http.StatusOK, gin.H{"message": "Order status deleted successfully"})
}
