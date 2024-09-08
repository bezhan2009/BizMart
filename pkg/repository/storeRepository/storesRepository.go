package storeRepository

import (
	"BizMart/db"
	"BizMart/logger"
	"BizMart/models"
	"BizMart/pkg/repository"
	"errors"
	"gorm.io/gorm"
)

// GetStores retrieves all stores without their hashed passwords.
func GetStores() ([]models.Store, error) {
	var stores []models.Store
	if err := db.GetDBConn().Omit("HashPassword").Find(&stores).Error; err != nil {
		logger.Error.Printf("[repository.GetStores] Error retrieving stores: %v", err)
		return nil, err
	}
	return stores, nil
}

// CreateStore adds a new store to the database.
func CreateStore(store *models.Store) error {
	if err := db.GetDBConn().Create(&store).Error; err != nil {
		logger.Error.Printf("[repository.CreateStore] Error creating store: %v", err)
		return err
	}
	return nil
}

// UpdateStore updates an existing store by ID.
func UpdateStore(storeID uint, updatedData *models.Store) error {
	// Start a transaction in case further operations are needed
	tx := db.GetDBConn().Begin()

	store, err := GetStoreByID(storeID)
	if err != nil {
		tx.Rollback() // rollback transaction on error
		return err
	}

	if err := tx.Model(&store).Updates(updatedData).Error; err != nil {
		logger.Error.Printf("[repository.UpdateStore] Error updating store with ID %d: %v", storeID, err)
		tx.Rollback()
		return err
	}

	tx.Commit() // commit the transaction
	return nil
}

// DeleteStore removes a store by ID from the database.
func DeleteStore(storeID uint) error {
	store, err := GetStoreByID(storeID)
	if err != nil {
		return err
	}

	if err := db.GetDBConn().Delete(&store).Error; err != nil {
		logger.Error.Printf("[repository.DeleteStore] Error deleting store with ID %d: %v", storeID, err)
		return err
	}
	return nil
}

// GetStoreByID retrieves a store by its ID.
func GetStoreByID(storeID uint) (models.Store, error) {
	var store models.Store
	if err := db.GetDBConn().Where("id = ?", storeID).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error.Printf("[repository.GetStoreByID] Store with ID %d not found: %v", storeID, err)
			return store, repository.TranslateGormError(err)
		}
		logger.Error.Printf("[repository.GetStoreByID] Error retrieving store with ID %d: %v", storeID, err)
		return store, repository.TranslateGormError(err)
	}
	return store, nil
}

func GetStoreByName(storeName string) (models.Store, error) {
	var store models.Store
	if err := db.GetDBConn().Where("name = ?", storeName).First(&store).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error.Printf("[repository.GetStoreByName] Store with Name %v not found: %v", storeName, err)
			return store, repository.TranslateGormError(err)
		}
		logger.Error.Printf("[repository.GetStoreByName] Error retrieving store with Name %v: %v", storeName, err)
		return store, repository.TranslateGormError(err)
	}
	return store, nil
}
