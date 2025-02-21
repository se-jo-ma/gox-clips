package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database_test() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
