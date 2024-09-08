package main

import (
	"BizMart/configs"
	"BizMart/internal/routes"
	security2 "BizMart/internal/security"
	db2 "BizMart/pkg/db"
	"BizMart/pkg/logger"
	"BizMart/server"
	"context"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var err error

func main() {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	err = godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("example.env")
		if err != nil {
			panic(errors.New(fmt.Sprintf("error loading .env file. Error is %s", err)))
		}
	}

	security2.AppSettings, err = configs.ReadSettings()
	if err != nil {
		panic(err)
	}
	security2.SetConnDB(security2.AppSettings)

	err = logger.Init()
	if err != nil {
		panic(err)
	}

	err = db2.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db2.Migrate()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	mainServer := new(server.Server)
	go func() {
		if err = mainServer.Run(security2.AppSettings.AppParams.PortRun, routes.InitRoutes(router)); err != nil {
			log.Fatalf("Ошибка при запуске HTTP сервера: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\n%s\n", yellow("Начало завершения сервиса"))

	if sqlDB, err := db2.GetDBConn().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Ошибка при закрытии соединения с БД: %s", err)
		}
	} else {
		log.Fatalf("Ошибка при получении *sql.DB из GORM: %s", err)
	}
	fmt.Println(green("Соединение с БД успешно закрыто"))

	if err = mainServer.Shutdown(context.Background()); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}

	fmt.Println(red("HTTP-сервис успешно выключен"))
	fmt.Println(red("Конец завершения программы"))
}
