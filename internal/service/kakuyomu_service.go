package service

import (
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
)

type KakuyomuService interface {
	KakuyomuListChapter(params *request.ChapterNovelRequest) (*model.BaseResponse[[]request.TranslateListRequest], error)
	KakuyomuChapterPage(params *request.ChapterNovelRequest) (*model.BaseResponse[*response.GetChapterPageResponse], error)
}
