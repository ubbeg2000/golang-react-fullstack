package api

import (
	"golang-api/api/routes"

	"github.com/gofiber/fiber/v2"
)

func New() {
	app := fiber.New()

	route := routes.New()

	route.UserRoutes.New(app)

	app.Listen(":3000")
}
