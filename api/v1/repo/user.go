package repo

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	conn *gorm.DB
}

// FindMany : find many users according to set paramateres
func (r *UserRepo) FindMany(limit int, page int) []models.User {
	var users []models.User
	r.conn.Limit(limit).Offset(limit * (page - 1)).Find(&users)
	return users
}

// FindOne : find one user accroding to set parameters
func (r *UserRepo) FindOne(param models.User) models.User {
	var user models.User
	r.conn.Where(&param).First(&user)
	return user
}

// Create : create new user
func (r *UserRepo) Create(param models.User) (models.User, error) {
	if res := r.conn.Create(&param); res.Error == nil {
		return param, nil
	} else {
		return param, res.Error
	}
}

// Update : update a user (user's ID is mandatory)
func (r *UserRepo) Update(param models.User) (models.User, error) {
	if res := r.conn.Model(&models.User{}).Updates(&param); res.Error == nil {
		return param, nil
	} else {
		return param, res.Error
	}
}

// Delete : delete one user
func (r *UserRepo) Delete(param models.User) error {
	if res := r.conn.Delete(&param); res.Error == nil {
		return nil
	} else {
		return res.Error
	}
}
