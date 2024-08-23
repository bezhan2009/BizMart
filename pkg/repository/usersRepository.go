package repository

import (
	"BizMart/db"
	"BizMart/logger"
	"BizMart/models"
	"gorm.io/gorm"
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

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.GetDBConn().Where("username = ?", username).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func UserExists(username, email string) (bool, bool, error) {
	users, err := GetAllUsers()
	if err != nil {
		return false, false, err
	}

	var usernameExists, emailExists bool
	for _, user := range users {
		if user.Username == username {
			usernameExists = true
		}
		if user.Email == email {
			emailExists = true
		}
	}
	return usernameExists, emailExists, nil
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

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, err
	}

	return user, nil
}
