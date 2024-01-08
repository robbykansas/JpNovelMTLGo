package main

import (
	"jpnovelmtlgo/internal/config"
	"jpnovelmtlgo/internal/controller"
	"jpnovelmtlgo/internal/exception"
	"jpnovelmtlgo/internal/repository"
	"jpnovelmtlgo/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New("application.yml")

	// Setup Repository
	TranslateRepository := repository.NewTranslateRepository(&configuration)

	// Setup Service
	SyosetuService := service.NewSyosetuService(&TranslateRepository)

	// Setup Controller
	HealthcheckController := controller.NewHealthcheckController(&configuration)
	SyosetuController := controller.NewSyosetuController(&SyosetuService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Helmet
	app.Use(helmet.New())

	// Setup Logging
	app.Use(logger.New())

	// Setup Routing
	HealthcheckController.Route(app)
	SyosetuController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
