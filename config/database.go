package config

import (
	"log"
	"os"

	"github.com/anything/smth/1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	err = DB.AutoMigrate(&models.Url{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
