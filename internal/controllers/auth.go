package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/pkg/errs"
	"BizMart/pkg/logger"
	utils2 "BizMart/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	var user models.User

	// Parse JSON body into the user struct
	if err := c.BindJSON(&user); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	// Check for missing fields
	if user.HashPassword == "" {
		HandleError(c, errs.ErrPasswordIsEmpty)
		return
	}

	if user.Email == "" {
		HandleError(c, errs.ErrEmailIsEmpty)
		return
	}

	if user.Username == "" {
		HandleError(c, errs.ErrUsernameIsEmpty)
		return
	}

	// Create the user
	userID, err := service.CreateUser(user)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrIncorrectUsernameOrPassword)
			return
		}

		HandleError(c, err)
		return
	}

	user.ID = userID

	// Generate access token
	accessToken, err := utils2.GenerateToken(user.ID, user.Username)
	if err != nil {
		logger.Error.Printf("Error generating access token: %s", err)
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"userID":       user.ID,
	})

	// Optionally, you can send a success message instead
	// c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func SignIn(c *gin.Context) {
	var user models.User

	// Parse JSON body into the user struct
	if err := c.BindJSON(&user); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	// Check for missing fields
	if user.HashPassword == "" {
		HandleError(c, errs.ErrPasswordIsEmpty)
		return
	}

	if user.Email == "" {
		HandleError(c, errs.ErrEmailIsEmpty)
		return
	}

	if user.Username == "" {
		HandleError(c, errs.ErrUsernameIsEmpty)
		return
	}

	// Hash the password before signing in
	user.HashPassword = utils2.GenerateHash(user.HashPassword)

	// Sign in the user
	user, accessToken, err := service.SignIn(user.Username, user.Email, user.HashPassword)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrIncorrectUsernameOrPassword.Error()})
			HandleError(c, errs.ErrIncorrectUsernameOrPassword)
			return
		}
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"user_id":      user.ID,
	})
}
