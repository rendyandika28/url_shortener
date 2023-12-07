package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"url_shortener/database"
	"url_shortener/helper"
	"url_shortener/router"
)

func Start() {
	database.OpenConnection()

	app := fiber.New(fiber.Config{
		ErrorHandler: helper.ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(logger.New())

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
