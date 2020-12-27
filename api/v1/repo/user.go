package repo

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	conn *gorm.DB
}

func (r *UserRepo) FindMany(limit int, page int) []models.User {
	var users []models.User
	r.conn.Limit(limit).Offset(limit * (page - 1)).Find(&users)
	return users
}

func (r *UserRepo) FindOne(param models.User) models.User {
	var user models.User
	r.conn.Where(&param).First(&user)
	return user
}
