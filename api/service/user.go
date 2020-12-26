package service

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/models"
	"github.com/ubbeg2000/golang-react-fullstack/api/repo"
)

type UserService struct {
	r repo.Repo
}

func (s *UserService) FindAll(limit int, page int) []models.User {
	return s.r.User.FindAll(limit, page)
}
