package main

import (
	"github.com/gin-gonic/gin"
	config2 "github.com/personal-finance/app/internal/config"
	incomeHandler "github.com/personal-finance/app/internal/delivery/http/income"
	incomeRepo "github.com/personal-finance/app/internal/repository/mysql/income"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var (
	db         = config2.InitPgsqlDB()
	repoIncome = incomeRepo.NewIncomeRepository(db)
)

func main() {
	cfg := config2.Server()

	app := gin.New()

	r := app.Group("/api")

	initHandler(cfg, r, repoIncome)

	err := app.Run(os.Getenv("APP_PORT"))
	if err != nil {
		log.Error("Can't Run App", zap.Error(err))
		os.Exit(0)
	}

}

func initHandler(cfg *config2.ServerConfig, r *gin.RouterGroup, incomeRepo incomeRepo.IncomeRepository) {
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello bang")
	})

	incomeHD := incomeHandler.NewHandler(cfg, incomeRepo)
	incomeRoute := r.Group("/income")
	incomeRoute.GET("/:userId", incomeHD.GetIncome)
	incomeRoute.GET("/:userId/:id", incomeHD.FindIncome)
	incomeRoute.POST("/:userId/add", incomeHD.AddIncome)
	incomeRoute.PUT("/:userId/:id", incomeHD.PutIncome)

}
