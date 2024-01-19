package test

import (
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"

	"github.com/gofiber/fiber/v2"
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

	resultMock := &response.ListChapterNovelResponse{
		StatusCode: "200",
		Data:       mockData,
	}

	uts.MockTranslateRepository.On("TranslateList", mock.AnythingOfType("[]request.TranslateListRequest")).Return(resultMock, nil)

	payload := &request.ChapterNovelRequest{
		Url: "https://kakuyomu.jp/works/16817139558533391541",
	}

	context := &fiber.Ctx{}

	res, err := uts.kakuyomuService.KakuyomuListChapter(context, payload)
	uts.Equal(resultMock, res)
	uts.Nil(err)
}

func (uts *UnitTestSuite) KakuyomuChapterPage() {
	mockTranslateResponse := &response.GetChapterPageResponse{
		Title:   "titleEn",
		Chapter: "chapterEn",
	}
	uts.MockTranslateRepository.On("TranslateChapter", mock.AnythingOfType("*request.TranslateChapterRequest")).Return(mockTranslateResponse, nil)

	payload := &request.ChapterNovelRequest{
		Url: "https://kakuyomu.jp/works/16817330664532961874/episodes/16817330664611957867",
	}
	context := &fiber.Ctx{}

	res, err := uts.kakuyomuService.KakuyomuChapterPage(context, payload)

	uts.Equal(mockTranslateResponse, res)
	uts.Nil(err)
}
