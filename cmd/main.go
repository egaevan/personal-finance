package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/personal-finance/domain/usecase"
	"github.com/personal-finance/internal/config"
	incomeHandler "github.com/personal-finance/internal/delivery/http/income"
	"github.com/personal-finance/internal/repository/mysql/income"
	incomeUC "github.com/personal-finance/internal/usecase/income"
	"log"
	"net/http"
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

	initHandler(e, useCaseIncome)
	http.Handle("/", e)

	//appLogger.Info("server start")
	e.GET("/", HealthCheck)
	e.Logger.Fatal(e.Start(":8000"))

}

func initHandler(e *echo.Echo, incomeUC usecase.IncomeUseCase) {
	incomeHandler.NewHandler(e, incomeUC)
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
