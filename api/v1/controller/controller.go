package controller

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/service"
)

var s service.Service = service.New()

type Controller struct {
	User UserController
}

func New() Controller {
	return Controller{
		UserController{s},
	}
}
