package main

import (
	"BizMart/db"
	"BizMart/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.Migrate()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	routes.SetupRouter(router)
}
