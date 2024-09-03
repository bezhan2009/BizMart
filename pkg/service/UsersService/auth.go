package UsersService

import (
	"BizMart/errs"
	"BizMart/models"
	"BizMart/pkg/repository/Users"
	"BizMart/utils"
)

func SignIn(username, useremail, password string) (user models.User, accessToken string, err error) {
	if useremail == "" && username == "" {
		return user, "", errs.ErrInvalidData
	}

	if useremail != "" && username != "" {
		user, err = Users.GetUserByEmailPasswordAndUsername(username, useremail, password)
		if err != nil {
			return user, "", errs.TranslateGormError(err)
		}
	} else if username != "" {
		user, err = Users.GetUserByUsernameAndPassword(username, password)
		if err != nil {
			return user, "", errs.TranslateGormError(err)
		}
	} else if useremail != "" {
		user, err = Users.GetUserByEmailAndPassword(useremail, password)
		if err != nil {
			return user, "", errs.TranslateGormError(err)
		}
	} else {
		return user, "", errs.ErrInvalidData
	}

	accessToken, err = utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return user, "", err
	}

	return user, accessToken, nil
}
