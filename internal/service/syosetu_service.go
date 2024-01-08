package service

import (
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"

	"github.com/gofiber/fiber/v2"
)

type SyosetuService interface {
	ListChapterNovel(ctx *fiber.Ctx, params *request.ListChapterNovelRequest) (*response.ListChapterNovelResponse, error)
	GetChapterPage(ctx *fiber.Ctx, params *request.ListChapterNovelRequest) (*response.GetChapterPageResponse, error)
}
