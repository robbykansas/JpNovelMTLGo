package controller

import (
	"jpnovelmtlgo/internal/config"

	"github.com/gofiber/fiber/v2"
)

type HealthcheckController struct {
	Configuration config.Config
}

func NewHealthcheckController(
	configuration *config.Config,
) HealthcheckController {
	return HealthcheckController{
		Configuration: *configuration,
	}
}

func (controller *HealthcheckController) Route(app *fiber.App) {
	app.Get("/health-check", controller.HealthCheck)
}

func (controller *HealthcheckController) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("API Running....")
}
