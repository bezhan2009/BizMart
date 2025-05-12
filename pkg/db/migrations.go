package db

import (
	models2 "BizMart/internal/app/models"
	"BizMart/internal/app/models/seeds"
	"BizMart/pkg/logger"
	"errors"
)

func Migrate() error {
	if dbConn == nil {
		logger.Error.Printf("[db.Migrate] Error because database connection is nil")

		return errors.New("database connection is not initialized")
	}

	//if userDBConn == nil {
	//	logger.Error.Printf("[db.Migrate] Error because users database connection is nil")
	//
	//	return errors.New("users database connection is not initialized")
	//}
	//
	//err := userDBConn.AutoMigrate(
	//	&models2.User{},
	//	&models2.Admin{},
	//)
	//if err != nil {
	//	logger.Error.Printf("[db.Migrate] Error migrating users tables: %v", err)
	//
	//	return err
	//}

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
		logger.Error.Printf("[db.Migrate] Error migrating tables: %v", err)

		return err
	}

	err = seeds.SeedCategories(dbConn)
	if err != nil {
		return err
	}

	err = seeds.SeedOrderStatuses(dbConn)
	if err != nil {
		return err
	}

	return nil
}
