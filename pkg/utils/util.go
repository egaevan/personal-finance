package utils

import (
	"os"
	"strconv"
)

type RestBody struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func ConvertBool(env string) bool {
	v, _ := strconv.ParseBool(os.Getenv(env))
	return v
}

func ConvertStrToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
