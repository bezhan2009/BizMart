package repository

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/logger"
)

func GetAllOrderByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := db.GetDBConn().
		Model(&models.Order{}).
		Where("user_id = ?", userID).
		Preload("OrderDetails").
		Find(&orders).Error; err != nil {
		logger.Error.Printf("[repository.GetAllOrderByUserID] Error getting orders by user id: %v", err)
		return []models.Order{}, TranslateGormError(err)
	}

	return orders, nil
}

func GetOrderByID(orderID uint) (models.Order, error) {
	var order models.Order
	if err := db.GetDBConn().
		Model(&models.Order{}).
		Where("id = ?", orderID).
		Preload("OrderDetails").
		First(&order).Error; err != nil {
		logger.Error.Printf("[repository.GetAllOrderByUserID] Error getting orders by user id: %v", err)
		return models.Order{}, TranslateGormError(err)
	}

	return order, nil
}

func GetOrderDetailsByID(orderDetailsID uint) (models.OrderDetails, error) {
	var orderDetails models.OrderDetails
	if err := db.GetDBConn().Model(models.OrderDetails{}).Where("id = ?", orderDetailsID).First(&orderDetails).Error; err != nil {
		logger.Error.Printf("[repository.GetOrderDetailsByID] Error getting orderDetails by id: %v", err)
		return models.OrderDetails{}, TranslateGormError(err)
	}

	return orderDetails, nil
}

func CreateOrder(order models.Order, orderDetails models.OrderDetails) error {
	if err := db.GetDBConn().Create(&order).Error; err != nil {
		logger.Error.Printf("[repository.CreateOrder] Error creating order: %v", err)
		return TranslateGormError(err)
	}

	if err := db.GetDBConn().Create(&orderDetails).Error; err != nil {
		logger.Error.Printf("[repository.CreateOrder] Error creating orderDetails: %v", err)
		return TranslateGormError(err)
	}

	return nil
}

func UpdateOrder(order models.Order, orderDetails models.OrderDetails) error {
	if err := db.GetDBConn().Save(&order).Error; err != nil {
		logger.Error.Printf("[repository.UpdateOrder] Error updating order: %v", err)
		return TranslateGormError(err)
	}

	if err := db.GetDBConn().Save(&orderDetails).Error; err != nil {
		logger.Error.Printf("[repository.UpdateOrder] Error updating orderDetails: %v", err)
		return TranslateGormError(err)
	}

	return nil
}

func DeleteOrder(orderID uint) error {
	order, err := GetOrderByID(orderID)
	if err != nil {
		return TranslateGormError(err)
	}

	if err := db.GetDBConn().Delete(&models.Order{}, order.ID).Error; err != nil {
		logger.Error.Printf("[repository.DeleteOrder] Error deleting order: %v", err)
		return TranslateGormError(err)
	}

	if err := db.GetDBConn().Delete(&models.OrderDetails{}, order.OrderDetailsID).Error; err != nil {
		logger.Error.Printf("[repository.DeleteOrder] Error deleting order: %v", err)
		return TranslateGormError(err)
	}

	return nil
}
