package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init(api fiber.Router) {
	api.Use(cors.New(cors.Config{
		AllowHeaders:  "Authorization,Content-Type,Upgrade,Accept,Acces-Control-Allow-Origin",
		ExposeHeaders: "Content-Type,Access-Control-Allow-Origin,Content-Type",
		MaxAge:        900,
	}))
	api.Use(csrf.New())
	api.Use(logger.New())
	api.Use(recover.New())
}
