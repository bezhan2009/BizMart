package middlewares

import (
	"BizMart/errs"
	"BizMart/pkg/repository"
	"BizMart/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func CheckAdmin(c *gin.Context) {
	authorization := c.Request.Header.Get(authorizationHeader)
	if authorization == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "No authorization header",
		})
		return
	}

	headerParts := strings.Split(authorization, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid auth header",
		})
		return
	}

	accessToken := headerParts[1]

	claims, err := utils.ParseToken(accessToken)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
	}

	user, err := repository.GetUserByID(claims.UserID)
	if err != nil {
		return
	}

	if user.ID == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "user not exist",
		})
		return
	}

	if user.Username != os.Getenv("ADMIN") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": errs.ErrPermissionDenied,
		})
		return
	}
	c.Next()
}
