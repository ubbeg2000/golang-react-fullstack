package api

import (
	v1 "github.com/ubbeg2000/golang-react-fullstack/api/v1"

	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/api")
	v1.Init(api)
}
