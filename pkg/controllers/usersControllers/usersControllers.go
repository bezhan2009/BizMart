package usersControllers

import (
	"BizMart/logger"
	"BizMart/models"
	"BizMart/pkg/service/UsersService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	users, err := UsersService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logger.Error.Printf("[controllers.GetAllUsers] error: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		logger.Error.Printf("[controllers.GetUserByID] invalid id: %s\n", c.Param("id"))
		return
	}

	user, err := UsersService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logger.Error.Printf("[controllers.GetUserByID] error: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Error.Printf("[controllers.CreateUser] error: %v\n", err)
		return
	}

	_, err := UsersService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logger.Error.Printf("[controllers.CreateUser] error: %v\n", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})
	logger.Info.Printf("[controllers.CreateUser] message successfully\n data %v", user)

}
