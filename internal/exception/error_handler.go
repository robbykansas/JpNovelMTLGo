package exception

import (
	"jpnovelmtlgo/internal/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//001 = Database
//002 = General
//003 = Session/Rdis

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, dataNotFoundError := err.(DataNotFoundError)
	if dataNotFoundError {
		status := fiber.StatusNotFound
		statusCodeString := strconv.Itoa(status)
		return ctx.Status(status).JSON(model.DefaultResponse{
			IsSuccessful: false,
			StatusCode:   statusCodeString,
			Message:      err.Error(),
		})
	}

	_, generalError := err.(GeneralError)
	if generalError {
		status := fiber.StatusBadRequest
		statusCodeString := strconv.Itoa(status)
		return ctx.Status(status).JSON(model.DefaultResponse{
			IsSuccessful: false,
			StatusCode:   statusCodeString,
			Message:      err.Error(),
		})
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		status := fiber.StatusUnauthorized
		statusCodeString := strconv.Itoa(status)
		return ctx.Status(status).JSON(model.DefaultResponse{
			IsSuccessful: false,
			StatusCode:   statusCodeString,
			Message:      err.Error(),
		})
	}

	_, internalServerError := err.(InternalServerError)
	if internalServerError {
		status := fiber.StatusInternalServerError
		statusCodeString := strconv.Itoa(status)
		return ctx.Status(status).JSON(model.DefaultResponse{
			IsSuccessful: false,
			StatusCode:   statusCodeString,
			Message:      err.Error(),
		})
	}

	return ctx.Status(400).JSON(model.DefaultResponse{
		IsSuccessful: false,
		StatusCode:   "400",
		Message:      "Please wait a few minutes before you try again.",
	})
}
