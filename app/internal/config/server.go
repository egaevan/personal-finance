package config

import (
	"github.com/personal-finance/pkg/utils"
	"os"
	"time"
)

type ServerConfig struct {
	Port    string
	Debug   bool
	Timeout time.Duration
}

func Server() *ServerConfig {
	return &ServerConfig{
		Port:    os.Getenv("APP_PORT"),
		Debug:   utils.ConvertBool("DEBUG"),
		Timeout: time.Duration(utils.ConvertInt("APP_TIMEOUT")) * time.Second,
	}
}
