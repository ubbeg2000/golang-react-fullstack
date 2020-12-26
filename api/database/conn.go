package database

import (
	models "github.com/ubbeg2000/goalng-react-fullstack/api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// New : create new *gorm.DB instance
func New() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./api/database/test.db"), &gorm.Config{})
	Migrate(db)
	if err != nil {
		panic("failed to connect database")
		return &gorm.DB{}, err
	}
	return db, nil
}

// Migrate : migrate existing existing models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Event{})
}
