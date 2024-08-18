package repository

import (
	"BizMart/db"
	"BizMart/logger"
	"BizMart/models"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %s\n", err.Error())
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		return user, err
	}
	return user, nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (err error) {
	//logger.Debug.Println(user.ID)
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return err
	}

	//logger.Debug.Println(user.ID)
	return nil
}
