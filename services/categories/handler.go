package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandlerGetAllCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := RepoGetAllCategories(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}
