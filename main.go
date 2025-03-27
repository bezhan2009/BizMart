package main

import (
	"BizMart/configs"
	ssogrpc "BizMart/internal/clients/sso/grpc"
	"BizMart/internal/jobs"
	"BizMart/internal/repository"
	"BizMart/internal/routes"
	security2 "BizMart/internal/security"
	"BizMart/internal/server"
	"BizMart/pkg/brokers/kafka"
	db2 "BizMart/pkg/db"
	"BizMart/pkg/logger"
	"context"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var err error

// @title BizMart API
// @version 1.3.2

// @description API Server for BizMart Application
// @host localhost:8585
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fmt.Println("STARTING BIZMART")
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

	err = db2.InitializeRedis(security2.AppSettings.RedisParams)
	if err != nil {
		panic(err)
	}

	err = db2.Migrate()
	if err != nil {
		panic(err)
	}

	err = kafka.CreateProducer(security2.AppSettings.KafkaParams)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	mainServer := new(server.Server)
	go func() {
		if err = mainServer.Run(security2.AppSettings.AppParams.PortRun, routes.InitRoutes(router)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error while starting HTTP Service: %s", err)
		}
	}()

	err = ssogrpc.New(
		context.Background(),
		security2.AppSettings.Clients.SSO.ClientAddress,
		security2.AppSettings.Clients.SSO.Timeout,
		security2.AppSettings.Clients.SSO.RetriesCount,
	)
	if err != nil {
		panic(err)
	}

	go repository.SynchronizationUserTable(security2.AppSettings.KafkaParams)

	go jobs.UpdateProductCache()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Printf("\n%s\n", yellow("Start of service termination"))

	// Закрытие соединения с БД
	err = db2.CloseDBConn()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error closing database connection: %s", err.Error()))
	}

	//err = db2.CloseUserDBConn()
	//if err != nil {
	//	fmt.Println(fmt.Sprintf("Error closing user database connection: %s", err.Error()))
	//}

	err = db2.CloseRedisConnection()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error closing redis connection: %s", err.Error()))
	}

	// Корректное завершение HTTP-сервера
	if err = mainServer.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error while termination HTTP Service: %s", err)
	} else {
		fmt.Println(green("HTTP-service termination successfully"))
	}

	fmt.Println(red("End of program completion"))
}
