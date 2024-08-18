package service

import (
	"BizMart/models"
	"BizMart/pkg/algorithms/PasswordAlgoritm"
	"BizMart/pkg/repository"
	"BizMart/utils"
	"errors"
	"gorm.io/gorm"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) error {
	_, err := repository.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	user.HashPassword = PasswordAlgoritm.Usage(user.HashPassword, true)
	user.HashPassword = utils.GenerateHash(user.HashPassword)

	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
