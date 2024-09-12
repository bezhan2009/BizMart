package repository

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/logger"
)

func GetAccountsByUserID(userID uint) ([]models.Account, error) {
	var accounts []models.Account
	if err := db.GetDBConn().Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
		logger.Error.Printf("[repository.GetMyAccounts] error getting accounts: %v\n", err)
		return nil, TranslateGormError(err)
	}

	return accounts, nil
}

func GetAccountByID(accountID uint) (models.Account, error) {
	var account models.Account
	if err := db.GetDBConn().Where("id = ?", accountID).First(&account).Error; err != nil {
		logger.Error.Printf("[repository.GetAccountByID] error getting account: %v\n", err)
		return models.Account{}, TranslateGormError(err)
	}

	return account, nil
}

func GetAccountByNumber(accountNumber string) (models.Account, error) {
	var account models.Account
	if err := db.GetDBConn().Where("account_number = ?", accountNumber).First(&account).Error; err != nil {
		logger.Error.Printf("[repository.GetAccountByNumber] error getting account: %v\n", err)
		return models.Account{}, TranslateGormError(err)
	}

	return account, nil
}

func FillAccountBalance(accountNumber string, amount float64) error {
	var account models.Account
	if err := db.GetDBConn().Where("account_number = ?", accountNumber).First(&account).Error; err != nil {
		logger.Error.Printf("[repository.GetAccountByNumber] error getting account: %v\n", err)
		return TranslateGormError(err)
	}

	account.Balance += amount

	if err := db.GetDBConn().Save(&account).Error; err != nil {
		logger.Error.Printf("[repository.GetAccountByNumber] error updating account: %v\n", err)
		return TranslateGormError(err)
	}

	return nil
}

func GetAccountByUserEmail(email string) (models.Account, error) {
	var account models.Account
	if err := db.GetDBConn().Where("user.email = ?", email).First(&account).Error; err != nil {
		logger.Error.Printf("[repository.GetAccountByUserEmail] error getting account: %v\n", err)
		return models.Account{}, TranslateGormError(err)
	}

	return account, nil
}

func CreateAccount(account models.Account) error {
	if err := db.GetDBConn().Create(&account).Error; err != nil {
		logger.Error.Printf("[repository.CreateAccount] error creating account: %v\n", err)
		return TranslateGormError(err)
	}

	return nil
}

func UpdateAccount(account models.Account) error {
	if err := db.GetDBConn().Save(&account).Error; err != nil {
		logger.Error.Printf("[repository.UpdateAccount] error updating account: %v\n", err)
		return TranslateGormError(err)
	}

	return nil
}

func DeleteAccount(account models.Account) error {
	if err := db.GetDBConn().Delete(&account).Error; err != nil {
		logger.Error.Printf("[repository.DeleteAccount] error deleting account: %v\n", err)
		return TranslateGormError(err)
	}

	return nil
}
