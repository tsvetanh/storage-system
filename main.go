package main

import (
	"log"
	"net/http"
	"storage/configuration"
	"storage/middleware"
	"storage/services/products"

	"github.com/gin-gonic/gin"
)

func main() {
	c := configuration.LoadConfig()

	r := gin.Default()
	r.Use(middleware.LoggingMiddleware)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "This is the version 1.0 of the backend app")
	})

	apiGroup := r.Group("/api")
	apiGroup.GET("/get-products", products.HandlerGetAllProducts(c))
	apiGroup.GET("/get-product", products.HandlerGetProductById(c))
	apiGroup.GET("/get-product-detailed", products.HandlerGetProductByIdDetailed(c))

	// Start the server
	if err := r.Run(":" + c.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
