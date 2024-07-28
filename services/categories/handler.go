package categories

import (
	"net/http"
	"storage/configuration"

	"github.com/gin-gonic/gin"
)

func HandlerGetAllCategories(conf *configuration.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := RepoGetAllCategories(conf.Db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}
