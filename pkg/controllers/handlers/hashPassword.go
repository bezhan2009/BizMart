package handlers

import (
	"BizMart/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

const SecretKeyHeader = "X-Secret-Key"

func HashPassword(c *gin.Context) {
	SecretKey := c.Request.Header.Get(SecretKeyHeader)
	if SecretKey == "" {
		c.JSON(400, gin.H{
			"error": "secret key is empty",
		})
		return
	}

	password := c.Query("password")
	password = strings.TrimSpace(password)
	password = utils.GenerateHash(password)
	c.JSON(200, gin.H{
		"password": password,
	})
}
