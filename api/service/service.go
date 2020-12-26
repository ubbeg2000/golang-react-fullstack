package service

import (
	"golang-api/api/repo"
)

var r repo.Repo = repo.New()

type Service struct {
	User  UserService
	Event EventService
}

func New() Service {
	return Service{
		UserService{r},
		EventService{r},
	}
}
