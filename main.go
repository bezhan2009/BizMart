package main

import (
	"BizMart/configs"
	"BizMart/db"
	"BizMart/logger"
	"BizMart/routes"
	"BizMart/security"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var err error

func main() {
	err = godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("example.env")
		if err != nil {
			panic(errors.New(fmt.Sprintf("error loading .env file. Error is %s", err)))
		}
	}

	security.AppSettings, err = configs.ReadSettings()
	if err != nil {
		panic(err)
	}
	security.SetConnDB(security.AppSettings)

	err = logger.Init()
	if err != nil {
		panic(err)
	}

	err = db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	routes.SetupRouter(router)
	err = router.Run(security.AppSettings.AppParams.PortRun)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := db.CloseDBConn()
		if err != nil {
			panic(err.Error())
		}
	}()
}
