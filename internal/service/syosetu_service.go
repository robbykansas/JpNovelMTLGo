package service

import (
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
)

type SyosetuService interface {
	ListChapterNovel(params *request.ChapterNovelRequest) (*model.BaseResponse[[]request.TranslateListRequest], error)
	GetChapterPage(params *request.ChapterNovelRequest) (*model.BaseResponse[*response.GetChapterPageResponse], error)
	JpEpub(params *request.ConvertNovelRequest) (*model.DefaultResponse, error)
	EnEpub(params *request.ConvertNovelRequest) (*model.DefaultResponse, error)
}
