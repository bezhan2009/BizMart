package middlewares

import (
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

const SecretKeyHeader = "X-Secret-Key"

func CheckSecretKey(c *gin.Context) {
	SecretKey := c.Request.Header.Get(SecretKeyHeader)
	if SecretKey == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "secret key header is empty",
		})
		return
	}

	if strings.TrimSpace(SecretKey) != os.Getenv("SECRET_KEY") {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "secret key does not match",
		})
		return
	}
	c.Next()
}
