package repository

import (
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
)

type TranslateRepository interface {
	TranslateChapter(params *request.TranslateChapterRequest) (*response.GetChapterPageResponse, error)
	TranslateList(params []request.TranslateListRequest) (*model.BaseResponse[[]request.TranslateListRequest], error)
	TranslateInfo(params *request.NovelInfo) (*response.TranslatedInfoResponse, error)
	TranslateListChapter(params []request.ChapterContent) ([]response.TranslatedListChapterResponse, error)
}
