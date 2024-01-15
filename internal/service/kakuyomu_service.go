package service

import (
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"

	"github.com/gofiber/fiber/v2"
)

type KakuyomuService interface {
	KakuyomuListChapter(ctx *fiber.Ctx, params *request.ChapterNovelRequest) (*response.ListChapterNovelResponse, error)
}
