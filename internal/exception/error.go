package exception

import (
	"jpnovelmtlgo/internal/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

func HandlerError(ctx *fiber.Ctx, err error) error {
	if ferr, ok := err.(*fiber.Error); ok {
		// Retrieve and use the status code
		status := ferr.Code
		statusCodeString := strconv.Itoa(status)
		return ctx.Status(status).JSON(model.DefaultResponse{
			IsSuccessful: false,
			StatusCode:   statusCodeString,
			Message:      err.Error(),
		})
	}

	status := fiber.StatusInternalServerError
	statusCodeString := strconv.Itoa(status)
	return ctx.Status(status).JSON(model.DefaultResponse{
		IsSuccessful: false,
		StatusCode:   statusCodeString,
		Message:      err.Error(),
	})
}
