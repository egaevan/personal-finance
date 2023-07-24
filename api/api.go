package api

import (
	"github.com/gin-gonic/gin"
	config2 "github.com/personal-finance/app/internal/config"
	incomeHandler "github.com/personal-finance/app/internal/delivery/http/income"
	incomeRepo "github.com/personal-finance/app/internal/repository/mysql/income"
	"net/http"
)

var (
	app        *gin.Engine
	db         = config2.InitPgsqlDB()
	repoIncome = incomeRepo.NewIncomeRepository(db)
)

// @schemes https http
// @host golang-vercel.vercel.app
func init() {
	cfg := config2.Server()

	app = gin.New()

	routes(cfg, app, repoIncome)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

func routes(cfg *config2.ServerConfig, app *gin.Engine, incomeRepo incomeRepo.IncomeRepository) {
	app.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello bang")
	})

	incomeHD := incomeHandler.NewHandler(cfg, incomeRepo)

	route := app.Group("/api")
	{
		incomeRoute := route.Group("/income")
		incomeRoute.GET("/:userId", incomeHD.GetIncome)
		incomeRoute.GET("/:userId/:id", incomeHD.FindIncome)
		incomeRoute.POST("/:userId/add", incomeHD.AddIncome)
		incomeRoute.PUT("/:userId/:id", incomeHD.PutIncome)
	}

}
