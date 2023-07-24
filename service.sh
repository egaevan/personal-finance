#!/bin/sh
export APP_TIMEOUT="2"
export DEBUG=true

export MYSQL_USERNAME="root"
export MYSQL_PASS="root"
export MYSQL_NAME="personal_finance"
export MYSQL_HOST="localhost"
export MYSQL_PORT="3306"

export PGSQL_USERNAME="default"
export PGSQL_PASS="yz98VfNIopJP"
export PGSQL_NAME="verceldb"
export PGSQL_HOST="ep-green-cloud-427063-pooler.ap-southeast-1.postgres.vercel-storage.com"
export PGSQL_PORT="5432"


export DB_TIMEZONE="Asia/Singapore"
export DB_MAX_CON_IDLE=10
export DB_MAX_CON_OPEN=100

go run cmd/main.go
# pause