package store

import (
	"BizMart/db"
	"BizMart/logger"
	"BizMart/models"
	"errors"
	"gorm.io/gorm"
)

// GetAllStoreReviews retrieves all reviews for a given store by its ID.
func GetAllStoreReviews(storeID uint) ([]models.StoreReview, error) {
	var storeReviews []models.StoreReview
	if err := db.GetDBConn().Where("store_id = ?", storeID).Find(&storeReviews).Error; err != nil {
		logger.Error.Printf("[repository.GetAllStoreReviews] Error retrieving store reviews for store ID %d: %v", storeID, err)
		return nil, err
	}

	return storeReviews, nil
}

// CreateStoreReview adds a new review for a store.
func CreateStoreReview(storeReview *models.StoreReview) error {
	if err := db.GetDBConn().Create(storeReview).Error; err != nil {
		logger.Error.Printf("[repository.CreateStoreReview] Error creating store review for store ID %d: %v", storeReview.StoreID, err)
		return err
	}
	return nil
}

// UpdateStoreReview updates an existing store review by its ID.
func UpdateStoreReview(storeReviewID uint, updatedData *models.StoreReview) error {
	// Start a transaction in case further operations are needed
	tx := db.GetDBConn().Begin()

	storeReview, err := GetStoreReviewByID(storeReviewID)
	if err != nil {
		tx.Rollback() // rollback transaction on error
		return err
	}

	if err := tx.Model(&storeReview).Updates(updatedData).Error; err != nil {
		logger.Error.Printf("[repository.UpdateStoreReview] Error updating store review with ID %d: %v", storeReviewID, err)
		tx.Rollback()
		return err
	}

	tx.Commit() // commit the transaction
	return nil
}

// DeleteStoreReview removes a store review by its ID.
func DeleteStoreReview(storeReviewID uint) error {
	storeReview, err := GetStoreReviewByID(storeReviewID)
	if err != nil {
		return err
	}

	if err := db.GetDBConn().Delete(&storeReview).Error; err != nil {
		logger.Error.Printf("[repository.DeleteStoreReview] Error deleting store review with ID %d: %v", storeReviewID, err)
		return err
	}

	return nil
}

// GetStoreReviewByID retrieves a store review by its ID.
func GetStoreReviewByID(storeReviewID uint) (models.StoreReview, error) {
	var storeReview models.StoreReview
	if err := db.GetDBConn().Where("id = ?", storeReviewID).First(&storeReview).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error.Printf("[repository.GetStoreReviewByID] Store review with ID %d not found: %v", storeReviewID, err)
			return storeReview, nil
		}
		logger.Error.Printf("[repository.GetStoreReviewByID] Error retrieving store review with ID %d: %v", storeReviewID, err)
		return storeReview, err
	}

	return storeReview, nil
}
