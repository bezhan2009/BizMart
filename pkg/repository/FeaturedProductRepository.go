package repository

import (
	"BizMart/db"
	"BizMart/logger"
	"BizMart/models"
)

// GetAllFeaturedProducts retrieves all featured products.
func GetAllFeaturedProducts() ([]models.FeaturedProduct, error) {
	var featuredProducts []models.FeaturedProduct

	if err := db.GetDBConn().Where("is_deleted = ?", false).Find(&featuredProducts).Error; err != nil {
		logger.Error.Printf("[repository.GetAllFeaturedProducts] Error retrieving featured products: %v\n", err)
		return nil, err
	}

	return featuredProducts, nil
}

// GetFeaturedProductByID retrieves a single featured product by its ID.
func GetFeaturedProductByID(featuredProductID uint) (models.FeaturedProduct, error) {
	var featuredProduct models.FeaturedProduct

	if err := db.GetDBConn().Where("id = ? AND is_deleted = ?", featuredProductID, false).First(&featuredProduct).Error; err != nil {
		logger.Error.Printf("[repository.GetFeaturedProductByID] Error retrieving featured product by ID: %v\n", err)
		return featuredProduct, err
	}

	return featuredProduct, nil
}

// CreateFeaturedProduct creates a new featured product.
func CreateFeaturedProduct(featuredProduct models.FeaturedProduct) error {
	if err := db.GetDBConn().Create(&featuredProduct).Error; err != nil {
		logger.Error.Printf("[repository.CreateFeaturedProduct] Error creating featured product: %v\n", err)
		return err
	}

	return nil
}

// UpdateFeaturedProduct updates an existing featured product.
func UpdateFeaturedProduct(featuredProductID uint, updatedData models.FeaturedProduct) error {
	featuredProduct, err := GetFeaturedProductByID(featuredProductID)
	if err != nil {
		return err
	}

	// Copy the updated data fields
	featuredProduct.ProductID = updatedData.ProductID
	featuredProduct.UserID = updatedData.UserID
	featuredProduct.IsDeleted = updatedData.IsDeleted

	if err := db.GetDBConn().Save(&featuredProduct).Error; err != nil {
		logger.Error.Printf("[repository.UpdateFeaturedProduct] Error updating featured product: %v\n", err)
		return err
	}

	return nil
}

// DeleteFeaturedProduct marks a featured product as deleted.
func DeleteFeaturedProduct(featuredProductID uint) error {
	featuredProduct, err := GetFeaturedProductByID(featuredProductID)
	if err != nil {
		return err
	}

	// Soft delete by marking as deleted
	featuredProduct.IsDeleted = true

	if err := db.GetDBConn().Save(&featuredProduct).Error; err != nil {
		logger.Error.Printf("[repository.DeleteFeaturedProduct] Error deleting featured product: %v\n", err)
		return err
	}

	return nil
}
