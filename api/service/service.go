package service

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/repo"
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
