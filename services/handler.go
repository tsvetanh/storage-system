package services

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandlerGetAllProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := RepoGetAllProducts(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}

}

func HandlerGetProductById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		productIdStr := c.Query("id")
		if productIdStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
			return
		}

		productId, err := strconv.ParseInt(productIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be a number: " + err.Error()})
			return
		}

		product, err := RepoGetProductById(db, productId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product: " + err.Error()})
			return
		}

		if product == (Product{}) {
			c.Status(http.StatusNoContent)
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func HandlerGetProductByIdDetailed(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		productIdStr := c.Query("id")
		if productIdStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
			return
		}

		productId, err := strconv.ParseInt(productIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be a number: " + err.Error()})
			return
		}

		product, err := RepoGetProductByIdDetailed(db, productId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product: " + err.Error()})
			return
		}

		if product == (Product{}) {
			c.Status(http.StatusNoContent)
			return
		}

		c.JSON(http.StatusOK, product)
	}
}
