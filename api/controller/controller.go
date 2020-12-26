package controller

import (
	"golang-api/api/service"
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
