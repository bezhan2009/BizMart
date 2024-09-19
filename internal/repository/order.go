package repository

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/logger"
)

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := db.GetDBConn().
		Model(&models.Order{}).
		Preload("OrderDetails").
		Find(&orders).Error; err != nil {
		logger.Error.Printf("[repository.GetAllOrderByUserID] Error getting orders by user id: %v", err)
		return []models.Order{}, TranslateGormError(err)
	}

	return orders, nil
}

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
	var product models.Product
	var err error

	if err = db.GetDBConn().Create(&orderDetails).Error; err != nil {
		logger.Error.Printf("[repository.CreateOrder] Error creating orderDetails: %v", err)
		return TranslateGormError(err)
	}

	order.OrderDetailsID = orderDetails.ID

	if err = db.GetDBConn().Create(&order).Error; err != nil {
		logger.Error.Printf("[repository.CreateOrder] Error creating order: %v", err)
		return TranslateGormError(err)
	}

	if product, err = GetProductByID(orderDetails.ProductID); err != nil {
		return TranslateGormError(err)
	}

	product.Amount -= orderDetails.Quantity

	if err = db.GetDBConn().Model(&models.Product{}).Where("id = ?", product.ID).Save(&product).Error; err != nil {
		logger.Error.Printf("[repository.CreateOrder] Error creating order: %v", err)
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

func GetNumberOfProductOrders(productID uint) (int, error) {
	orders, err := GetAllOrders()
	if err != nil {
		return 0, TranslateGormError(err)
	}

	numOfProductOrders := 0

	for _, order := range orders {
		if order.OrderDetails.ProductID == productID {
			numOfProductOrders++
		}
	}

	return numOfProductOrders, nil
}

func GetNumberOfStoreOrders(storeID uint) (int, error) {
	products, err := GetProductByStoreIDWithoutFilters(storeID)
	if err != nil {
		return 0, TranslateGormError(err)
	}

	numOfStoreOrders := 0

	for _, product := range products {
		if product.StoreID == storeID {
			numOfProductOrders, err := GetNumberOfProductOrders(product.ID)
			if err != nil {
				return 0, TranslateGormError(err)
			}

			numOfStoreOrders += numOfProductOrders
		}
	}

	return numOfStoreOrders, nil
}

func GetNumberOfStoreProducts(storeID uint) (int, error) {
	products, err := GetProductByStoreIDWithoutFilters(storeID)
	if err != nil {
		return 0, TranslateGormError(err)
	}

	numOfProductOrders := 0

	for _, product := range products {
		if product.StoreID == storeID {
			numOfProductOrders++
		}
	}

	return numOfProductOrders, nil
}
