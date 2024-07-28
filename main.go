package main

import (
	"log"
	"net/http"
	"storage/configuration"
	"storage/middleware"
	"storage/services/categories"
	"storage/services/product"

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
	apiGroup.GET("/get-products", product.HandlerGetAllProducts(db))
	apiGroup.GET("/get-product", product.HandlerGetProductById(db))
	apiGroup.GET("/get-product-detailed", product.HandlerGetProductByIdDetailed(db))
	apiGroup.GET("/get-categories", categories.HandlerGetAllCategories(db))

	// Start the server
	if err := r.Run(":" + c.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
