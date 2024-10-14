package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL driver
)

var DB *gorm.DB

func ConnectDatabase(dbUrl string) {
	var err error
	DB, err = gorm.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto-migrate models
	DB.AutoMigrate(&User{}, &Geospatial{}) // Ensure models are defined
}
