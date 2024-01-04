package main

import (
	"jpnovelmtlgo/internal/config"
	"jpnovelmtlgo/internal/controller"
	"jpnovelmtlgo/internal/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New("application.yml")

	// Setup Controller
	HealthcheckController := controller.NewHealthcheckController(&configuration)

	// Setup Repository

	// Setup Service

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Helmet
	app.Use(helmet.New())

	// Setup Logging
	app.Use(logger.New())

	// Setup Routing
	HealthcheckController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
