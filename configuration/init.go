package configuration

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"os"
)

var Configuration *Config
var Database *gorm.DB

func LoadConfig() *Config {
	data, err := os.ReadFile("./configuration/config.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	err = json.Unmarshal(data, &Configuration)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return Configuration
}

func SetUpDatabase(conf *Config) *gorm.DB {
	dsn := "user=" + conf.Database.User +
		" password=" + conf.Database.Password +
		" dbname=" + conf.Database.DatabaseName +
		" port=" + conf.Database.Port +
		" host=" + conf.Database.Host +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
