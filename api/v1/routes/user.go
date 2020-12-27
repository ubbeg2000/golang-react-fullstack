package routes

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/controller"

	"github.com/gofiber/fiber/v2"
)

var c controller.Controller = controller.New()

func UserRoutes(app fiber.Router) {
	app.Group("/user", c.User.FindAll)
}
