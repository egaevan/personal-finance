package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/personal-finance/pkg/config"
	"net/url"
	"os"
	"strings"
)

func InitMysqlDB() (*sql.DB, error) {
	cfg := config.DatabaseMySQL()
	var (
		errMysql error
		dbConn   *sql.DB
	)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", cfg.TimeZone)
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	if os.Getenv("APP_ENV") == "production" {
		dbConn, errMysql = sql.Open(`nrmysql`, dsn)
	} else {
		dbConn, errMysql = sql.Open(`mysql`, dsn)
	}

	if errMysql != nil {
		return nil, errMysql
	}

	dbConn.SetMaxOpenConns(cfg.MaxConnectionOpen)
	dbConn.SetMaxIdleConns(cfg.MaxConnectionIdle)
	dbConn.SetConnMaxLifetime(cfg.TimeOut)

	return dbConn, nil
}

func InitPgsqlDB() (*sql.DB, error) {
	cfg := config.DatabasePgSQL()
	configs := []string{
		"host=" + cfg.Host,
		"user=" + cfg.User,
		"password=" + cfg.Password,
		"dbname=" + cfg.Database,
		"port=" + cfg.Port,
		"TimeZone=" + cfg.TimeZone,
		"sslmode=require",
		"search_path=" + cfg.Schema,
	}
	dsn := strings.Join(configs, " ")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
