package database

import (
	models "golang-api/api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./api/database/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		return &gorm.DB{}, err
	} else {
		return db, nil
	}
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Event{})
}
