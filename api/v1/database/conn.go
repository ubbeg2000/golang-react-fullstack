package database

import (
	models "github.com/ubbeg2000/golang-react-fullstack/api/v1/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// New : create new *gorm.DB instance
func New() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./api/v1/database/test.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	Migrate(db)

	if err != nil {
		panic("failed to connect database")
		return &gorm.DB{}, err
	}
	return db, nil
}

// Migrate : migrate existing existing models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Location{},
		&models.Comment{},
		&models.Event{},
		&models.Group{},
		&models.Role{},
		&models.User{},
		&models.Link{},
		&models.Tag{},
	)
}
