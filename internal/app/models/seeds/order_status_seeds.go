package seeds

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/logger"
	"errors"
	"gorm.io/gorm"
)

func SeedOrderStatuses(db *gorm.DB) error {
	// Определяем стандартные категории
	orderStatuses := []models.OrderStatus{
		{StatusName: "New", Description: "New order"},
		{StatusName: "Pending", Description: "User waiting for his order"},
		{StatusName: "Payed", Description: "User has paid for order"},
		{StatusName: "Received", Description: "User has received order"},
	}

	for _, orderStatus := range orderStatuses {
		// Проверяем, существует ли категория
		var existingOrderStatus models.OrderStatus
		if err := db.First(&existingOrderStatus, "status_name = ?", orderStatus.StatusName).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Если категория не найдена, создаем её
				db.Create(&orderStatus)
			} else {
				// Обработка других ошибок
				logger.Error.Printf("[seeds.SeedOrderStatuses] Error seeding orderStatuses: %v", err)

				return err
			}
		}
	}

	return nil
}
