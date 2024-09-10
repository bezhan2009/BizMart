package db

import (
	"BizMart/internal/security"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func ConnectToDB() error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		security.HostName,
		security.Port,
		security.UserName,
		security.Password,
		security.DBName,
		security.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	dbConn = db
	return nil
}

func CloseDBConn() error {
	//err := dbConn.Close()
	//if err != nil {
	//	return err
	//}

	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
