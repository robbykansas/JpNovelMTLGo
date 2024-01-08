package controller

import (
	"jpnovelmtlgo/internal/exception"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/service"

	"github.com/gofiber/fiber/v2"
)

type SyosetuController struct {
	SyosetuService service.SyosetuService
}

func NewSyosetuController(
	SyosetuService *service.SyosetuService,
) SyosetuController {
	return SyosetuController{
		SyosetuService: *SyosetuService,
	}
}

func (controller *SyosetuController) Route(app *fiber.App) {
	app.Post("/syosetu/list-chapter", controller.ListChapterNovel)
	app.Post("/syosetu/chapter", controller.GetChapterNovel)
}

func (controller *SyosetuController) ListChapterNovel(ctx *fiber.Ctx) error {
	var request request.ListChapterNovelRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.SyosetuService.ListChapterNovel(ctx, &request)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(200).JSON(response)
}

func (controller *SyosetuController) GetChapterNovel(ctx *fiber.Ctx) error {
	var request request.ListChapterNovelRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.SyosetuService.GetChapterPage(ctx, &request)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(200).JSON(response)
}
