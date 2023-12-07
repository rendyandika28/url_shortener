package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"url_shortener/controller"
	"url_shortener/database"
	"url_shortener/repository"
	"url_shortener/service"
)

func RegisterShortenRoutes(app fiber.Router, validate *validator.Validate) {
	// Link shortener app
	shortenerRepository := repository.NewShortenerRepository()
	shortenerService := service.NewShortenerService(shortenerRepository, database.DB, validate)
	shortenerController := controller.NewShortenerController(shortenerService)

	shortener := app.Group("/shortener")
	shortener.Get("/:slug", shortenerController.FindBySlug)
	shortener.Delete("/:slug", shortenerController.Delete)
	shortener.Put("/:slug", shortenerController.Update)
	shortener.Get("/", shortenerController.FindAll)
	shortener.Post("/", shortenerController.Create)
}
