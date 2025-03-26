package models

import "time"

type Configs struct {
	LogParams      LogParams      `json:"log_params"`
	AppParams      AppParams      `json:"app_params"`
	PostgresParams PostgresParams `json:"postgres_params"`
	RedisParams    RedisParams    `json:"redis_params"`
	KafkaParams    KafkaParams    `json:"kafka_params"`
	Clients        ClientsConfig  `json:"clients"`
	Auth           Auth           `json:"auth"`
}

type LogParams struct {
	LogDirectory     string `json:"log_directory"`
	LogInfo          string `json:"log_info"`
	LogError         string `json:"log_error"`
	LogWarn          string `json:"log_warn"`
	LogDebug         string `json:"log_debug"`
	MaxSizeMegabytes int    `json:"max_size_megabytes"`
	MaxBackups       int    `json:"max_backups"`
	MaxAge           int    `json:"max_age"`
	Compress         bool   `json:"compress"`
	LocalTime        bool   `json:"local_time"`
}

type AppParams struct {
	ServerURL  string `json:"server_url"`
	ServerName string `json:"server_name"`
	AppVersion string `json:"app_version"`
	PortRun    string `json:"port_run"`
	GinMode    string `json:"gin_mode"`
	Env        string `json:"env"`
}

type PostgresParams struct {
	User         string `json:"user"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	Database     string `json:"database"`
	UserDatabase string `json:"user_database"`
	SSLMode      string `json:"sslmode"`
}

type RedisParams struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type KafkaParams struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Topic           string `json:"topic"`
	GroupID         string `json:"group_id"`
	AutoOffsetReset string `json:"auto_offset_reset"`
}

type Auth struct {
	JwtSecretKey  string        `json:"jwt_secret_key"`
	JwtTtlMinutes time.Duration `json:"jwt_ttl_minutes"`
}

type Client struct {
	ClientAddress string        `json:"address"`
	Timeout       time.Duration `json:"timeout"`
	RetriesCount  int           `json:"retries_count"`
	Insecure      bool          `json:"insecure"`
}

type ClientsConfig struct {
	SSO Client `json:"sso"`
}
