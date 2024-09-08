package orderRepository

import (
	"BizMart/db"
	"BizMart/errs"
	"BizMart/logger"
	"BizMart/models"
	"BizMart/pkg/repository"
)

func GetOrderStatusByID(orderStatusID uint) (models.OrderStatus, error) {
	var orderStatus models.OrderStatus
	if err := db.GetDBConn().Where("id = ?", orderStatusID).First(&orderStatus).Error; err != nil {
		logger.Error.Printf("[repository.GetOrderStatusByID] error getting order status by ID: %s\n", err.Error())
		return orderStatus, repository.TranslateGormError(err)
	}

	return orderStatus, nil
}

func GetOrderStatusByName(orderStatusName string) (models.OrderStatus, error) {
	var orderStatus models.OrderStatus
	if err := db.GetDBConn().Where("status_name = ?", orderStatusName).First(&orderStatus).Error; err != nil {
		logger.Error.Printf("[repository.GetOrderStatusByID] error getting order status by ID: %s\n", err.Error())
		return orderStatus, repository.TranslateGormError(err)
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

func UpdateOrderStatus(orderStatusID uint, orderStatus models.OrderStatus) (OrderStatusID uint, err error) {
	existingOrderStatus := models.Category{}
	if err = db.GetDBConn().First(&existingOrderStatus, orderStatusID).Error; err != nil {
		logger.Error.Printf("[repository.UpdateOrderStatus] orderStatus not found: %v\n", err)
		return 0, errs.ErrOrderStatusNotFound
	}

	if err = db.GetDBConn().Model(&existingOrderStatus).Updates(orderStatus).Error; err != nil {
		logger.Error.Printf("[repository.UpdateOrderStatus] error updating orderStatus: %v\n", err)
		return orderStatusID, err
	}

	return orderStatusID, nil
}

func DeleteOrderStatus(orderStatus models.OrderStatus) error {
	if err := db.GetDBConn().Delete(&orderStatus).Error; err != nil {
		logger.Error.Printf("[repository.DeleteOrderStatus] error deleting order status: %s\n", err.Error())
		return err
	}

	return nil
}
