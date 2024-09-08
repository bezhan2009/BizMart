package repository

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/logger"
)

func GetMyAccounts(userID uint) ([]models.Account, error) {
	var accounts []models.Account
	if err := db.GetDBConn().Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
		logger.Error.Printf("[repository.GetMyAccounts] error getting accounts: %v\n", err)
		return nil, err
	}

	return accounts, nil
}

func GetAccountByID(accountID uint) (models.Account, error) {
	var account models.Account
	if err := db.GetDBConn().Where("account_id = ?", accountID).First(&account).Error; err != nil {
		logger.Error.Printf("[repository.GetAccountByID] error getting account: %v\n", err)
		return models.Account{}, err
	}

	return account, nil
}

func GetAccountByUserEmail(email string) (models.Account, error) {
	var account models.Account
	if err := db.GetDBConn().Where("user.email = ?", email).First(&account).Error; err != nil {
		logger.Error.Printf("[repository.GetAccountByUserEmail] error getting account: %v\n", err)
		return models.Account{}, err
	}

	return account, nil
}

func CreateAccount(account models.Account) (uint, error) {
	if err := db.GetDBConn().Create(&account).Error; err != nil {
		logger.Error.Printf("[repository.CreateAccount] error creating account: %v\n", err)
		return 0, err
	}

	return account.ID, nil
}

func UpdateAccount(account models.Account) (uint, error) {
	if err := db.GetDBConn().Save(&account).Error; err != nil {
		logger.Error.Printf("[repository.UpdateAccount] error updating account: %v\n", err)
		return 0, err
	}

	return account.ID, nil
}

func DeleteAccount(account models.Account) (uint, error) {
	if err := db.GetDBConn().Delete(&account).Error; err != nil {
		logger.Error.Printf("[repository.DeleteAccount] error deleting account: %v\n", err)
		return 0, err
	}

	return account.ID, nil
}
