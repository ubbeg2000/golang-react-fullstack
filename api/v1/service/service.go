package service

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/repo"
)

var r repo.Repo = repo.New()

type Service struct {
	User  UserService
	Event EventService
	Auth  AuthService
}

var (
	userService  UserService  = UserService{r}
	eventService EventService = EventService{r}
	authService  AuthService  = AuthService{r}
)

func New() Service {
	return Service{
		userService,
		eventService,
		authService,
	}
}
