package db

import (
	models2 "BizMart/internal/app/models"
	"errors"
)

func Migrate() error {
	if dbConn == nil {
		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(
		&models2.User{},
		&models2.Store{},
		&models2.StoreReview{},
		&models2.Address{},
		&models2.UserProfile{},
		&models2.Account{},
		&models2.Category{},
		&models2.Comment{},
		&models2.FeaturedProduct{},
		&models2.Product{},
		&models2.ProductImage{},
		&models2.Order{},
		&models2.OrderDetails{},
		&models2.OrderStatus{},
		&models2.Review{},
		&models2.Payment{},
	)

	if err != nil {
		return err
	}

	return nil
}
