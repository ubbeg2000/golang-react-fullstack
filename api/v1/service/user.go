package service

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/helper/types"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/models"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/repo"
)

type UserService struct {
	repo repo.Repo
}

func (s *UserService) FindAll(q *types.UserQuery) []models.User {
	return s.repo.User.FindMany(q.Limit, q.Page)
}
