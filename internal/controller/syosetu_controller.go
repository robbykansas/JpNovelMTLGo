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
	SyosetuService service.SyosetuService,
) SyosetuController {
	return SyosetuController{
		SyosetuService: SyosetuService,
	}
}

func (controller *SyosetuController) Route(app *fiber.App) {
	app.Post("/syosetu/list-chapter", controller.ListChapterNovel)
	app.Post("/syosetu/chapter", controller.GetChapterNovel)
	app.Post("/syosetu/jp-epub", controller.JpEpub)
	app.Post("/syosetu/en-epub", controller.EnEpub)
}

func (controller *SyosetuController) ListChapterNovel(ctx *fiber.Ctx) error {
	var request request.ChapterNovelRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return exception.HandlerError(ctx, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	response, err := controller.SyosetuService.ListChapterNovel(&request)
	if err != nil {
		return exception.HandlerError(ctx, err)
	}

	return ctx.Status(200).JSON(response)
}

func (controller *SyosetuController) GetChapterNovel(ctx *fiber.Ctx) error {
	var request request.ChapterNovelRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return exception.HandlerError(ctx, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	response, err := controller.SyosetuService.GetChapterPage(&request)
	if err != nil {
		return exception.HandlerError(ctx, err)
	}

	return ctx.Status(200).JSON(response)
}

func (controller *SyosetuController) JpEpub(ctx *fiber.Ctx) error {
	var request request.ConvertNovelRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return exception.HandlerError(ctx, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	response, err := controller.SyosetuService.JpEpub(&request)
	if err != nil {
		return exception.HandlerError(ctx, err)
	}

	return ctx.Status(200).JSON(response)
}

func (controller *SyosetuController) EnEpub(ctx *fiber.Ctx) error {
	var request request.ConvertNovelRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return exception.HandlerError(ctx, fiber.NewError(fiber.StatusBadRequest, err.Error()))
	}

	response, err := controller.SyosetuService.EnEpub(&request)
	if err != nil {
		return exception.HandlerError(ctx, err)
	}

	return ctx.Status(200).JSON(response)
}
