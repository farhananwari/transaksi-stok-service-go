package routes

import (
	"net/http"
	"transaksi-stok-service-go/controller"
	"transaksi-stok-service-go/database"
	"transaksi-stok-service-go/repositories"
	"transaksi-stok-service-go/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine) {

	db := database.GetDB()

	// Initialize repositories

	stokBalancesRepo := repositories.NewStokBalancesRepository(db)
	masterDataRepo := repositories.NewMasterDataRepository(db)
	// Initialize services
	stokBalancesService := services.NewStokBalancesService(stokBalancesRepo, masterDataRepo)
	// Initialize controller
	stokBalancesRepoController := controller.NewStokBalancesController(stokBalancesService)
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "transaksi-stok-service-go",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Root endpoint
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "transaksi-stokService API v1",
				"version": "1.0.0",
			})
		})
		v1.GET("/stok-balances/:locationID", stokBalancesRepoController.GetBarangWithStok)
		v1.POST("/stok-balances", stokBalancesRepoController.CreateStokBalance)
	}
}
