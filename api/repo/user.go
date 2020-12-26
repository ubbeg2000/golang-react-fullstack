package repo

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	conn *gorm.DB
}

func (r *UserRepo) FindAll(limit int, page int) []models.User {
	var users []models.User
	r.conn.Limit(limit).Offset(limit * (page - 1)).Find(&users)
	return users
}
