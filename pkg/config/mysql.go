package config

import (
	"github.com/personal-finance/pkg/utils"
	"os"
	"time"
)

type DatabaseMySQLConfig struct {
	User              string
	Password          string
	Database          string
	Host              string
	Port              string
	TimeZone          string
	MaxConnectionIdle int
	MaxConnectionOpen int
	Schema            string
	Debug             bool
	TimeOut           time.Duration
}

func DatabaseMySQL() DatabaseMySQLConfig {
	schema := "public"
	cfgSchema := os.Getenv("DB_SCHEMA")
	if cfgSchema != "" {
		schema = cfgSchema
	}

	return DatabaseMySQLConfig{
		User:              os.Getenv("MYSQL_USERNAME"),
		Password:          os.Getenv("MYSQL_PASS"),
		Database:          os.Getenv("MYSQL_NAME"),
		Host:              os.Getenv("MYSQL_HOST"),
		Port:              os.Getenv("MYSQL_PORT"),
		TimeZone:          os.Getenv("DB_TIMEZONE"),
		MaxConnectionIdle: utils.ConvertInt("DB_MAX_CON_IDLE"),
		MaxConnectionOpen: utils.ConvertInt("DB_MAX_CON_OPEN"),
		Schema:            schema,
		Debug:             utils.ConvertBool("DEBUG"),
		TimeOut:           time.Duration(utils.ConvertInt("APP_TIMEOUT")) * time.Second,
	}
}
