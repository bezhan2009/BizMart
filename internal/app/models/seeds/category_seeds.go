package seeds

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/logger"
	"errors"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) error {
	// Определяем стандартные категории
	categories := []models.Category{
		{CategoryName: "Books", Description: "A wide range of books from various genres."},
		{CategoryName: "Electronics", Description: "Latest gadgets and electronic devices."},
		{CategoryName: "Clothing", Description: "Fashionable clothing for all ages."},
		{CategoryName: "Home & Kitchen", Description: "Essentials for your home and kitchen."},
		{CategoryName: "Beauty", Description: "Beauty products and personal care items."},
	}

	for _, category := range categories {
		// Проверяем, существует ли категория
		var existingCategory models.Category
		if err := db.First(&existingCategory, "category_name = ?", category.CategoryName).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Если категория не найдена, создаем её
				db.Create(&category)
			} else {
				// Обработка других ошибок
				logger.Error.Printf("[seeds.SeedCategories] Error seeding categories: %v", err)

				return err
			}
		}
	}

	return nil
}
