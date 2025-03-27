package grpc

import (
	"BizMart/internal/app/models"
	"BizMart/internal/security"
	"BizMart/pkg/logger"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	ssov1 "github.com/bezhan2009/AuthProtos/gen/go/sso"
	"os"
	"strconv"
)

func encryptSecret(secret string) (string, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("KEY")))
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12) // 12 байт для GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	encrypted := aesGCM.Seal(nil, nonce, []byte(secret), nil)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (c *Client) SignUp(ctx context.Context, user models.User) (int64, error) {
	const op = "grpc.SignUp"

	resp, err := c.authApi.Register(ctx, &ssov1.RegisterRequest{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.HashPassword,
	})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.GetUserId(), nil
}

func (c *Client) SignIn(ctx context.Context, user models.User) (models.TokenResponse, error) {
	const op = "grpc.SignIn"

	appId := security.AppSettings.AppParams.AppID
	var err error
	if appId == 0 {
		appId, err = strconv.Atoi(os.Getenv("APP_ID"))
		if err != nil {
			logger.Error.Printf("[%s]: %s]", op, err.Error())

			return models.TokenResponse{}, nil
		}
	}
	resp, err := c.authApi.Login(ctx, &ssov1.LoginRequest{
		Username: user.Username,
		Password: user.HashPassword,
		AppLogin: int32(appId),
	})
	if err != nil {
		logger.Error.Printf("[%s]: %s", op, err)

		return models.TokenResponse{}, fmt.Errorf("%s: %w", op, err)
	}

	return models.TokenResponse{
		AccessToken:  resp.GetAccessToken(),
		RefreshToken: resp.GetRefreshToken(),
		UserID:       uint(resp.GetUserId()),
	}, nil
}

func (c *Client) IsAdmin(ctx context.Context, userID int32) (bool, error) {
	const op = "grpc.IsAdmin"

	resp, err := c.authApi.IsAdmin(ctx, &ssov1.IsAdminRequest{
		UserId: userID,
	})
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return resp.IsAdmin, nil
}

func (c *Client) Ping(ctx context.Context, message string) (*ssov1.PingResponse, error) {
	const op = "grpc.Ping"

	resp, err := c.pingApi.Ping(ctx, &ssov1.PingRequest{
		Message: message,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp, nil
}
