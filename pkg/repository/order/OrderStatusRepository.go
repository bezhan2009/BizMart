package order

import (
	"BizMart/db"
	"BizMart/logger"
	"BizMart/models"
)

func GetOrderStatusByID(orderStatusID uint) (models.OrderStatus, error) {
	var orderStatus models.OrderStatus
	if err := db.GetDBConn().Where("id = ?", orderStatusID).First(&orderStatus).Error; err != nil {
		logger.Error.Printf("[repository.GetOrderStatusByID] error getting order status by ID: %s\n", err.Error())
		return orderStatus, err
	}

	return orderStatus, nil
}

func GetOrderStatusByName(orderStatusName string) (models.OrderStatus, error) {
	var orderStatus models.OrderStatus
	if err := db.GetDBConn().Where("status_name = ?", orderStatusName).First(&orderStatus).Error; err != nil {
		logger.Error.Printf("[repository.GetOrderStatusByID] error getting order status by ID: %s\n", err.Error())
		return orderStatus, err
	}

	return orderStatus, nil
}

func GetAllOrderStatuses() ([]models.OrderStatus, error) {
	var orderStatuses []models.OrderStatus
	if err := db.GetDBConn().Find(&orderStatuses).Error; err != nil {
		logger.Error.Printf("[repository.GetAllOrderStatuses] error getting all order statuses: %s\n", err.Error())
		return orderStatuses, err
	}

	return orderStatuses, nil
}

func CreateOrderStatus(orderStatus models.OrderStatus) (uint, error) {
	if err := db.GetDBConn().Create(&orderStatus).Error; err != nil {
		logger.Error.Printf("[repository.CreateOrderStatus] error creating order status: %s\n", err.Error())
		return 0, err
	}

	return orderStatus.ID, nil
}

func UpdateOrderStatus(orderStatus models.OrderStatus) (uint, error) {
	if err := db.GetDBConn().Save(&orderStatus).Error; err != nil {
		logger.Error.Printf("[repository.UpdateOrderStatus] error updating order status: %s\n", err.Error())
		return 0, err
	}

	return orderStatus.ID, nil
}

func DeleteOrderStatus(orderStatus models.OrderStatus) error {
	if err := db.GetDBConn().Delete(&orderStatus).Error; err != nil {
		logger.Error.Printf("[repository.DeleteOrderStatus] error deleting order status: %s\n", err.Error())
		return err
	}

	return nil
}
