package controller

import (
	"github.com/gofiber/fiber/v2"
	"url_shortener/model/web"
)

func ApiCheck(c *fiber.Ctx) error {
	return c.JSON(web.Response{
		Code:   200,
		Status: "Success",
	})
}
