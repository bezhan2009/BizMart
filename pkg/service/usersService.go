package service

import (
	"BizMart/models"
	"BizMart/pkg/algorithms/PasswordAlgoritm"
	"BizMart/pkg/repository"
	"BizMart/validators"
	"fmt"
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
	usernameExists, emailExists, err := repository.UserExists(user.Username, user.Email)
	if err != nil {
		return fmt.Errorf("failed to check existing user: %w", err)
	}

	if user.HashPassword == "" || user.Email == "" || user.Username == "" {
		return fmt.Errorf("invalid data")
	}

	isValid := validators.Password(user.HashPassword)

	if !isValid {
		return fmt.Errorf("invalid password")
	}
	user.HashPassword = PasswordAlgoritm.Usage(user.HashPassword, true)

	if usernameExists {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}

	if emailExists {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	if err := repository.CreateUser(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
