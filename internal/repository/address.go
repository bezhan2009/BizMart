package repository

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/logger"
)

// GetMyAddresses retrieves all addresses for a given user.
func GetMyAddresses(userID uint) ([]models.Address, error) {
	var addresses []models.Address
	if err := db.GetDBConn().Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
		logger.Error.Printf("[repository.GetMyAddresses] error getting user addresses: %v\n", err)
		return nil, TranslateGormError(err)
	}
	return addresses, nil
}

// GetAddressByID retrieves an address by its ID.
func GetAddressByID(addressID uint) (*models.Address, error) {
	var address models.Address
	if err := db.GetDBConn().Where("id = ?", addressID).First(&address).Error; err != nil {
		logger.Error.Printf("[repository.GetAddressByID] error getting address by ID: %v\n", err)
		return nil, TranslateGormError(err)
	}
	return &address, nil
}

// GetAddressByNameAndUserID retrieves an address by its Address name.
func GetAddressByNameAndUserID(addressName string, userID uint) (*models.Address, error) {
	var address models.Address
	if err := db.GetDBConn().Where("address_name = ? AND user_id = ?", addressName, userID).First(&address).Error; err != nil {
		logger.Error.Printf("[repository.GetAddressByNameAndUserID] error getting address by Address name: %v\n", err)
		return nil, TranslateGormError(err)
	}
	return &address, nil
}

// CreateAddress creates a new address and returns its ID.
func CreateAddress(address *models.Address) error {
	if err := db.GetDBConn().Create(address).Error; err != nil {
		logger.Error.Printf("[repository.CreateAddress] error creating address: %v\n", err)
		return TranslateGormError(err)
	}
	return nil
}

// UpdateAddress updates an existing address and returns its ID.
func UpdateAddress(address *models.Address) error {
	if err := db.GetDBConn().Save(address).Error; err != nil {
		logger.Error.Printf("[repository.UpdateAddress] error updating address: %v\n", err)
		return TranslateGormError(err)
	}
	return nil
}

// DeleteAddress deletes an address by ID and returns its ID.
func DeleteAddress(addressID uint) error {
	if err := db.GetDBConn().Delete(&models.Address{}, addressID).Error; err != nil {
		logger.Error.Printf("[repository.DeleteAddress] error deleting address: %v\n", err)
		return TranslateGormError(err)
	}
	return nil
}
