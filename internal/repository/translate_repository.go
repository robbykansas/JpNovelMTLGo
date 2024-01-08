package repository

import (
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
)

type TranslateRepository interface {
	TranslateChapter(params *request.TranslateChapterRequest) (*response.GetChapterPageResponse, error)
}
