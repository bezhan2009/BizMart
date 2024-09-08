package middlewares

import (
	"BizMart/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	UserIDCtx           = "userID"
)

func CheckUserAuthentication(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is empty",
		})
		return
	}

	accessToken := headerParts[1]

	claims, err := utils.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(claims)

	c.Set(UserIDCtx, claims.UserID)
	c.Next()
}
