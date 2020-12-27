package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/ubbeg2000/golang-react-fullstack/api"
	"github.com/ubbeg2000/golang-react-fullstack/config"
)

func main() {
	config.New()

	app := fiber.New()

	api.Init(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name": "Grame",
			"age":  20,
		})
	})

	log.Fatal(app.Listen(viper.GetString(`server.port`)))
}
