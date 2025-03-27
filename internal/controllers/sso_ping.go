package controllers

import (
	ssogrpc "BizMart/internal/clients/sso/grpc"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SSOPing(c *gin.Context) {
	resp, err := ssogrpc.GetClient().Ping(context.Background(), "ping")
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp.GetReply()})
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "pong"})
}
