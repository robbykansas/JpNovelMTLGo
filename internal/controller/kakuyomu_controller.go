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
	KakuyomuService service.KakuyomuService,
) KakuyomuController {
	return KakuyomuController{
		KakuyomuService: KakuyomuService,
	}
}

func (controller *KakuyomuController) Route(app *fiber.App) {
	app.Post("/kakuyomu/list-chapter", controller.KakuyomuListChapter)
	app.Post("/kakuyomu/chapter", controller.KakuyomuChapterPage)
}

func (controller *KakuyomuController) KakuyomuListChapter(ctx *fiber.Ctx) error {
	var request request.ChapterNovelRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	response, err := controller.KakuyomuService.KakuyomuListChapter(&request)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return ctx.Status(200).JSON(response)
}

func (controller *KakuyomuController) KakuyomuChapterPage(ctx *fiber.Ctx) error {
	var request request.ChapterNovelRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		// Generate this error with remove header content-type
		exception.PanicIfNeeded(err)
	}

	response, err := controller.KakuyomuService.KakuyomuChapterPage(&request)
	if err != nil {
		// fmt.Println(err.(*fiber.Error).Code)
		return exception.HandlerError(ctx, err)
	}

	return ctx.Status(200).JSON(response)
}
