package routes

import (
	"golang-api/api/controller"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct {
	c controller.Controller
}

func (r *UserRoutes) New(app *fiber.App) {
	app.Group("/user", r.c.User.FindAll)
}
