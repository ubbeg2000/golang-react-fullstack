package controller

import (
	"strconv"

	"github.com/ubbeg2000/golang-react-fullstack/api/v1/helper/types"

	"github.com/ubbeg2000/golang-react-fullstack/api/v1/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service service.Service
}

func (c *UserController) FindAll(ctx *fiber.Ctx) error {
	q := types.DefaultUserQuery()

	if err := ctx.QueryParser(q); err == nil {
		data := c.service.User.FindAll(q)
		return ctx.JSON(fiber.Map{
			"data": data,
			"msg":  "On page " + strconv.Itoa(q.Page) + " with limit " + strconv.Itoa(q.Limit),
		})
	} else {
		ctx.SendStatus(400)
		return ctx.JSON(fiber.Map{
			"code": 400,
			"msg":  err.Error(),
		})
	}

	return nil
}
