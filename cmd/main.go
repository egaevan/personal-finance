package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/personal-finance/internal/config"
	"github.com/personal-finance/internal/repository/mysql/income"
	incomeUC "github.com/personal-finance/internal/usecase/income"
	"log"
)

func main() {
	e := echo.New()

	cfg := config.Server()

	db, err := config.InitMysqlDB()
	if err != nil {
		fmt.Println("Error connect to db")
		log.Fatal(err)
	}

	defer db.Close()

	// Init Repository
	repoInconme := income.NewIncomeRepository(db)

	// Init UseCase
	useCaseIncome := incomeUC.NewIncomeInteractor(cfg, repoInconme)

}

func initHandler(e *echo.Echo, cfg *config.ServerConfig) {

}
