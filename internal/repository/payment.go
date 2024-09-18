package repository

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/db"
	"BizMart/pkg/logger"
)

func GetAllUserPayments(userID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := db.GetDBConn().Model(models.Payment{}).Where("user_id = ?", userID).Find(&payments).Error; err != nil {
		logger.Error.Printf("[repository.GetAllUserPayments] error getting user payments: %v\n", err)
		return []models.Payment{}, TranslateGormError(err)
	}

	return payments, nil
}

func GetPaymentByID(paymentID uint) (models.Payment, error) {
	var payment models.Payment
	if err := db.GetDBConn().Model(models.Payment{}).Where("id = ?", paymentID).First(&payment).Error; err != nil {
		logger.Error.Printf("[repository.GetPaymentByID] error getting payment By ID: %s\n", err.Error())
		return models.Payment{}, TranslateGormError(err)
	}

	return payment, nil
}

func CreatePayment(payment models.Payment) error {
	if err := db.GetDBConn().Model(models.Payment{}).Create(&payment).Error; err != nil {
		logger.Error.Printf("[repository.CreatePayment] error creating new payment: %s\n", err.Error())
		return TranslateGormError(err)
	}

	return nil
}

func UpdatePayment(payment models.Payment) error {
	if err := db.GetDBConn().Model(models.Payment{}).Where("id = ?", payment.ID).Save(&payment).Error; err != nil {
		logger.Error.Printf("[repository.UpdatePayment] error updating payment: %s\n]", err.Error())
		return TranslateGormError(err)
	}

	return nil
}

func DeletePayment(payment models.Payment) error {
	if err := db.GetDBConn().Model(models.Payment{}).Delete(&payment).Error; err != nil {
		logger.Error.Printf("[repository.DeletePayment] error deleting payment: %s\n", err.Error())
		return TranslateGormError(err)
	}

	return nil
}
