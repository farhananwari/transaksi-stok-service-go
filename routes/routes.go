package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine) {

	//db := database.GetDB()

	// Initialize repositories

	// Initialize services

	// Initialize controller

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
	}

}
