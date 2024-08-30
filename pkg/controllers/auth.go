package controllers

import (
	"BizMart/errs"
	"BizMart/logger"
	"BizMart/models"
	"BizMart/pkg/service"
	"BizMart/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
	var user models.User

	// Parse JSON body into the user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check for missing fields
	if user.HashPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrPasswordIsEmpty.Error()})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrEmailIsEmpty.Error()})
		return
	}

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrUsernameIsEmpty.Error()})
		return
	}

	// Create the user
	userID, err := service.CreateUser(user)
	if err != nil {
		handleError(c, err)
		return
	}

	user.ID = userID

	// Generate access token
	accessToken, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		logger.Error.Printf("Error generating access token: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
		handleError(c, err)
		return
	}

	// Check for missing fields
	if user.HashPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrPasswordIsEmpty.Error()})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrEmailIsEmpty.Error()})
		return
	}

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrUsernameIsEmpty.Error()})
		return
	}

	// Hash the password before signing in
	user.HashPassword = utils.GenerateHash(user.HashPassword)

	// Sign in the user
	user, accessToken, err := service.SignIn(user.Username, user.Email, user.HashPassword)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"user_id":      user.ID,
	})
}
