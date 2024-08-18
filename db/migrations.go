package db

import "BizMart/models"

func Migrate() error {
	err := dbConn.AutoMigrate(models.User{},
		models.Store{},
		models.StoreReview{})
	if err != nil {
		return err
	}
	return nil
}
