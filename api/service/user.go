package service

import (
	"golang-api/api/models"
	"golang-api/api/repo"
)

type UserService struct {
	r repo.Repo
}

func (s *UserService) FindAll(limit int, page int) []models.User {
	return s.r.User.FindAll(limit, page)
}
