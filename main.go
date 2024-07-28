package main

import (
	"log"
	"net/http"
	"storage/configuration"
	"storage/middleware"
	"storage/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	c := configuration.LoadConfig()

	// Initialize database connection
	db := configuration.SetUpDatabase(c)

	// Set up Gin router
	r := gin.Default()
	r.Use(middleware.LoggingMiddleware)

	// Define routes
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "This is the version 1.0 of the backend app")
	})

	apiGroup := r.Group("/api")
	apiGroup.GET("/get-products", services.HandlerGetAllProducts(db))
	apiGroup.GET("/get-product", services.HandlerGetProductById(db))
	apiGroup.GET("/get-product-detailed", services.HandlerGetProductByIdDetailed(db))

	// Start the server
	if err := r.Run(":" + c.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
