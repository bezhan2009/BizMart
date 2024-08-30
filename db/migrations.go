package db

import (
	"BizMart/models"
	"errors"
)

func Migrate() error {
	if dbConn == nil {
		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(
		&models.User{},
		&models.Store{},
		&models.StoreReview{},
		&models.Address{},
		&models.UserProfile{},
		&models.Account{},
		&models.Category{},
		&models.Comment{},
		&models.FeaturedProduct{},
		&models.Product{},
		&models.ProductImage{},
		&models.Order{},
		&models.OrderDetails{},
		&models.OrderStatus{},
		&models.Review{},
		&models.Payment{},
	)

	if err != nil {
		return err
	}

	return nil
}
