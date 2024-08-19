package main

import (
	"BizMart/db"
	"BizMart/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	routes.SetupRouter(router)
	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
