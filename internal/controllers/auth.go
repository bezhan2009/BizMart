package controllers

import (
	"BizMart/internal/app/models"
	"BizMart/internal/app/service"
	"BizMart/pkg/errs"
	"BizMart/pkg/logger"
	utils2 "BizMart/pkg/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

// SignUp godoc
// @Summary Register a new user
// @Description This endpoint registers a new user with a username, email, and password.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.UserRequest true "User information"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/sign-up [post]
func SignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

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

	accessToken, refreshToken, err := utils2.GenerateToken(user.ID, user.Username)
	if err != nil {
		logger.Error.Printf("Error generating access token: %s", err)
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID,
	})
}

// SignIn godoc
// @Summary User login
// @Description This endpoint logs in an existing user using their username, email, and password.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.UserLogin true "User login information"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/sign-in [post]
func SignIn(c *gin.Context) {
	var user models.User
	isEmailEmpty := false

	if err := c.BindJSON(&user); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if user.HashPassword == "" {
		HandleError(c, errs.ErrPasswordIsEmpty)
		return
	}

	if user.Email == "" {
		isEmailEmpty = true
	}

	if user.Username == "" && isEmailEmpty {
		HandleError(c, errs.ErrUsernameIsEmpty)
		return
	}

	user.HashPassword = utils2.GenerateHash(user.HashPassword)

	user, accessToken, refreshToken, err := service.SignIn(user.Username, user.Email, user.HashPassword)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			HandleError(c, errs.ErrIncorrectUsernameOrPassword)
			return
		}

		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID,
	})
}

// RefreshToken godoc
// @Summary Refresh Token
// @Description This endpoint refreshes the access token.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.RefreshRequest true "User login information"
// @Success 200 {object} models.RefreshTokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/refresh [post]
func RefreshToken(c *gin.Context) {
	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	// Проверка валидности refresh_token
	token, err := jwt.ParseWithClaims(requestBody.RefreshToken, &utils2.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		HandleError(c, errs.ErrInvalidToken)
		return
	}

	// Генерация нового access_token
	claims, ok := token.Claims.(*utils2.CustomClaims)
	if !ok || claims.ExpiresAt < time.Now().Unix() {
		HandleError(c, errs.ErrRefreshTokenExpired)
		return
	}

	accessToken, _, err := utils2.GenerateToken(claims.UserID, claims.Username)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
