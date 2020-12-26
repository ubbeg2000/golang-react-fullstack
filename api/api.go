package api

import (
	"github.com/ubbeg2000/golang-react-fullstack/api/routes"

	"github.com/gofiber/fiber/v2"
)

func New() {
	app := fiber.New()

	route := routes.New()

	route.UserRoutes.New(app)

	app.Listen(":3000")
}
