package controller

import "github.com/gofiber/fiber/v2"

type ShortenerController interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindBySlug(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}
