package controllers

import (
	"BizMart/logger"
	"BizMart/models"
	"BizMart/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllUsers(c *gin.Context) {
	logger.Info.Printf("Client with ip: [%s] requested list of users\n", c.ClientIP())
	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
	logger.Info.Printf("Client with ip: [%s] got list of users\n", c.ClientIP())
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("[controllers.GetUserByID] invalid user_id path parameter: %s\n", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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

		return
	}

	err := service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})

}
