package test

import (
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"

	"github.com/stretchr/testify/mock"
)

func (uts *UnitTestSuite) TestKakuyomuListChapter() {
	mockData := []request.TranslateListRequest{
		{
			Title:   "example1",
			Url:     "url1",
			TitleEn: "exampleEn",
			Order:   1,
		},
		{
			Title:   "example2",
			Url:     "url2",
			TitleEn: "exampleEn2",
			Order:   2,
		},
	}

	mockResult := &model.BaseResponse[[]request.TranslateListRequest]{
		StatusCode: "200",
		Message:    "Success",
		Data:       mockData,
	}

	uts.MockTranslateRepository.On("TranslateList", mock.AnythingOfType("[]request.TranslateListRequest")).Return(mockResult, nil)

	payload := &request.ChapterNovelRequest{
		Url: "https://kakuyomu.jp/works/16817139558533391541",
	}

	res, err := uts.kakuyomuService.KakuyomuListChapter(payload)
	uts.Equal(mockResult, res)
	uts.Nil(err)
}

func (uts *UnitTestSuite) TestKakuyomuChapterPage() {
	mockTranslateResponse := &response.GetChapterPageResponse{
		Title:   "titleEn",
		Chapter: "chapterEn",
	}
	uts.MockTranslateRepository.On("TranslateChapter", mock.AnythingOfType("*request.TranslateChapterRequest")).Return(mockTranslateResponse, nil)

	payload := &request.ChapterNovelRequest{
		Url: "https://kakuyomu.jp/works/16817330664532961874/episodes/16817330664611957867",
	}

	mockResult := &model.BaseResponse[*response.GetChapterPageResponse]{
		StatusCode: "200",
		Message:    "Success",
		Data:       mockTranslateResponse,
	}

	res, err := uts.kakuyomuService.KakuyomuChapterPage(payload)

	uts.Equal(mockResult, res)
	uts.Nil(err)
}
