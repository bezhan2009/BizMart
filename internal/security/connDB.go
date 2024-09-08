package security

import (
	"BizMart/internal/app/models"
	"os"
)

var (
	HostName string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
)

func SetConnDB(AppSettingsConfig models.Configs) {
	AppSettings = AppSettingsConfig
	var postgresParams = AppSettings.PostgresParams
	HostName = postgresParams.Host
	Port = postgresParams.Port
	UserName = postgresParams.User
	Password = os.Getenv("DB_PASSWORD")
	DBName = postgresParams.Database
	SSLMode = postgresParams.SSLMode
}
