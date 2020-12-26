package controller

import (
	"golang-api/api/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	s service.Service
}

type qstring struct {
	Limit int `query:"limit"`
	Page  int `query:"page"`
}

func (c *UserController) FindAll(ctx *fiber.Ctx) error {
	q := new(qstring)

	if err := ctx.QueryParser(q); err == nil {
		data := c.s.User.FindAll(q.Limit, q.Page)
		return ctx.JSON(data)
	}

	return nil
}
