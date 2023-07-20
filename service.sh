#!/bin/sh
export APP_TIMEOUT="2"
export DEBUG=true

export MYSQL_USERNAME="root"
export MYSQL_PASS="root"
export MYSQL_NAME="personal_finance"
export MYSQL_HOST="localhost"
export MYSQL_PORT="3306"
export DB_TIMEZONE="Asia/Jakarta"
export DB_MAX_CON_IDLE=10
export DB_MAX_CON_OPEN=100

go run cmd/main.go
# pause