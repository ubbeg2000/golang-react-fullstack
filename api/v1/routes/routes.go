package routes

import (
	"github.com/gofiber/fiber/v2"
)

func New(api fiber.Router) {
	UserRoutes(api)
}
