package routes

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/controller"
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
