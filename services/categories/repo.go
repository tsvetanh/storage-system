package categories

import (
	// "database/sql"

	"gorm.io/gorm"
)

func RepoGetAllCategories(db *gorm.DB) ([]Categories, error) {
	var categories []Categories
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
