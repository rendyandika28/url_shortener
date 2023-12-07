package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"url_shortener/controller"
	"url_shortener/helper"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	validate := validator.New()
	if err := validate.RegisterValidation("slug", helper.ValidationSlug); err != nil {
		return
	}

	api := app.Group("/api", logger.New())
	api.Get("/", controller.ApiCheck)

	// Link shortener App
	RegisterShortenRoutes(api, validate)
}
