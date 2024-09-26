package utils

import (
	"BizMart/internal/security"
	"BizMart/pkg/errs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// CustomClaims определяет кастомные поля токена
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken генерирует JWT токен с кастомными полями
func GenerateToken(userID uint, username string) (string, string, error) {
	// Access token
	claims := &CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(), // токен истекает через 1 час
			Issuer:    security.AppSettings.AppParams.ServerName,
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", "", err
	}

	// Refresh token
	refreshClaims := &CustomClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // токен истекает через 72 часа
			Issuer:    security.AppSettings.AppParams.ServerName,
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// ParseToken парсит JWT токен и возвращает кастомные поля
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errs.ErrInvalidToken
}
