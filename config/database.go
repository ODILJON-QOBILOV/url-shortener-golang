package config

import (
	"log"

	"github.com/anything/smth/1/models"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	_ "modernc.org/sqlite"  // <-- pure Go SQLite driver
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("file:mydb.db?cache=shared&mode=rwc"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.Url{})
}
