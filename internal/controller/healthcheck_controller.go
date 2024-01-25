package controller

import (
	"github.com/gofiber/fiber/v2"
)

type HealthcheckController struct{}

func NewHealthcheckController() HealthcheckController {
	return HealthcheckController{}
}

func (controller *HealthcheckController) Route(app *fiber.App) {
	app.Get("/health-check", controller.HealthCheck)
}

func (controller *HealthcheckController) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("API Running....")
}
