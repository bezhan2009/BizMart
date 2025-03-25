package db

import (
	"BizMart/internal/security"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	dbConn     *gorm.DB
	userDBConn *gorm.DB
)

func ConnectToDB() error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		security.HostName,
		security.Port,
		security.UserName,
		security.Password,
		security.DBName,
		security.SSLMode,
	)

	//connUserStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	//	security.HostName,
	//	security.Port,
	//	security.UserName,
	//	security.Password,
	//	security.UserDBName,
	//	security.SSLMode,
	//)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	dbConn = db
	//
	//userDB, err := gorm.Open(postgres.Open(connUserStr), &gorm.Config{})
	//if err != nil {
	//	return err
	//}
	//
	//userDBConn = userDB

	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}

//func GetUserDBConn() *gorm.DB {
//	return userDBConn
//}

func CloseDBConn() error {
	if sqlDB, err := GetDBConn().DB(); err == nil {
		if err = sqlDB.Close(); err != nil {
			log.Fatalf("Error while closing DB: %s", err)
		}
		fmt.Println("Connection closed successfully")
	} else {
		log.Fatalf("Error while getting *sql.DB from GORM: %s", err)
	}

	return nil
}

//func CloseUserDBConn() error {
//	if sqlDB, err := GetUserDBConn().DB(); err == nil {
//		if err = sqlDB.Close(); err != nil {
//			log.Fatalf("Error while closing user DB: %s", err)
//		}
//		fmt.Println("Connection closed successfully")
//	} else {
//		log.Fatalf("Error while getting *sql.DB from GORM: %s", err)
//	}
//
//	return nil
//}
