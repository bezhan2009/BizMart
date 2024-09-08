package repository

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/errs"
	"BizMart/pkg/logger"
)

func GetAllCategories() (categories []models.Category, err error) {
	if err = db.GetDBConn().Find(&categories).Error; err != nil {
		logger.Error.Printf("[repository.GetAllCategories] error finding all categories: %v", err)
		return nil, err
	}

	return categories, nil
}

func GetCategoryByID(categoryID uint) (category models.Category, err error) {
	if err = db.GetDBConn().Where("id = ?", categoryID).First(&category).Error; err != nil {
		logger.Error.Printf("[repository.GetCategoryByID] error getting category by id: %v, err: %v", categoryID, err)
		return category, err
	}

	return category, nil
}

func GetCategoryByName(categoryName string) (category models.Category, err error) {
	if err = db.GetDBConn().Where("category_name = ?", categoryName).First(&category).Error; err != nil {
		logger.Error.Printf("[repository.GetCategoryByName] error getting category by name: %s\n", err)
		return category, err
	}

	return category, nil
}

func CreateCategory(category models.Category) (categoryID uint, err error) {
	if err = db.GetDBConn().Create(&category).Error; err != nil {
		logger.Error.Printf("[repository.CreateCategory] error creating category: %v\n", err)
		return category.ID, err
	}

	return category.ID, nil
}

func UpdateCategory(categID uint, category models.Category) (categoryID uint, err error) {
	existingCategory := models.Category{}
	if err = db.GetDBConn().First(&existingCategory, categID).Error; err != nil {
		logger.Error.Printf("[repository.UpdateCategory] category not found: %v\n", err)
		return 0, errs.ErrCategoryNotFound
	}

	if err = db.GetDBConn().Model(&existingCategory).Updates(category).Error; err != nil {
		logger.Error.Printf("[repository.UpdateCategory] error updating category: %v\n", err)
		return categID, err
	}

	return categID, nil
}

func DeleteCategory(categoryID uint) (err error) {
	var category models.Category
	category, err = GetCategoryByID(categoryID)
	if err != nil {
		return errs.ErrCategoryNotFound
	}

	if err = db.GetDBConn().Delete(&category).Error; err != nil {
		logger.Error.Printf("[repository.DeleteCategory] error deleting category: %v\n", err)
		return err
	}

	return nil
}
