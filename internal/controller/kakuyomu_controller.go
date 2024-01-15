package controller

import (
	"jpnovelmtlgo/internal/exception"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/service"

	"github.com/gofiber/fiber/v2"
)

type KakuyomuController struct {
	KakuyomuService service.KakuyomuService
}

func NewKakuyomuController(
	KakuyomuService *service.KakuyomuService,
) KakuyomuController {
	return KakuyomuController{
		KakuyomuService: *KakuyomuService,
	}
}

func (controller *KakuyomuController) Route(app *fiber.App) {
	app.Post("/kakuyomu/list-chapter", controller.KakuyomuListChapter)
}

func (controller *KakuyomuController) KakuyomuListChapter(ctx *fiber.Ctx) error {
	var request request.ChapterNovelRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.KakuyomuService.KakuyomuListChapter(ctx, &request)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(200).JSON(response)
}
