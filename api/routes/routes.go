package routes

import (
	"golang-api/api/controller"
)

var c controller.Controller = controller.New()

type Routes struct {
	UserRoutes UserRoutes
}

func New() Routes {
	return Routes{
		UserRoutes{c},
	}
}
