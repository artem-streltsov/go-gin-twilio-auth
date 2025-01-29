package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
