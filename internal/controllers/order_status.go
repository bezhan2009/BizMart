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
