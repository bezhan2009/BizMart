package service

import (
	"BizMart/internal/app/models"
	"BizMart/internal/repository"
	"BizMart/pkg/errs"
	"BizMart/pkg/utils"
)

func SignIn(username, useremail, password string) (user models.User, accessToken string, refreshToken string, err error) {
	if useremail == "" && username == "" {
		return user, "", "", errs.ErrInvalidData
	}

	if useremail != "" && username != "" {
		user, err = repository.GetUserByEmailPasswordAndUsername(username, useremail, password)
		if err != nil {
			return user, "", "", repository.TranslateGormError(err)
		}
	} else if username != "" {
		user, err = repository.GetUserByUsernameAndPassword(username, password)
		if err != nil {
			return user, "", "", repository.TranslateGormError(err)
		}
	} else if useremail != "" {
		user, err = repository.GetUserByEmailAndPassword(useremail, password)
		if err != nil {
			return user, "", "", repository.TranslateGormError(err)
		}
	} else {
		return user, "", "", errs.ErrInvalidData
	}

	accessToken, refreshToken, err = utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return user, "", "", err
	}

	return user, accessToken, refreshToken, nil
}
