package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/middleware"
	"github.com/ubbeg2000/golang-react-fullstack/api/v1/routes"
)

func Init(api fiber.Router) {
	v1 := api.Group("/v1")
	middleware.Init(v1)
	routes.New(v1)
}
