package test

import (
	"jpnovelmtlgo/internal/model"
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

	mockResult := &model.BaseResponse[[]request.TranslateListRequest]{
		StatusCode: "200",
		Message:    "Success",
		Data:       mockData,
	}

	uts.MockTranslateRepository.On(
		"TranslateList",
		mock.AnythingOfType("[]request.TranslateListRequest")).
		Return(mockResult, nil)

	payload := &request.ChapterNovelRequest{
		Url: "https://kakuyomu.jp/works/16817139558533391541",
	}

	res, err := uts.kakuyomuService.KakuyomuListChapter(payload)
	uts.Equal(mockResult, res)
	uts.Nil(err)

	uts.MockTranslateRepository.AssertExpectations(uts.T())
}

func (uts *UnitTestSuite) TestKakuyomuChapterPage() {
	mockTranslateResponse := &response.GetChapterPageResponse{
		Title:   "titleEn",
		Chapter: "chapterEn",
	}
	uts.MockTranslateRepository.On(
		"TranslateChapter",
		mock.AnythingOfType("*request.TranslateChapterRequest")).
		Return(mockTranslateResponse, nil).Once()

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

	uts.MockTranslateRepository.AssertExpectations(uts.T())
}

// Still error because of request mock.AnythingOfType, the result is like the past result because of same request
func (uts *UnitTestSuite) TestKakuyomuChapterPage_Error() {
	errData := fiber.NewError(fiber.StatusBadGateway, "Bad Gateway Error")
	uts.MockTranslateRepository.On(
		"TranslateChapter",
		mock.AnythingOfType("*request.TranslateChapterRequest")).
		Return(nil, errData).Once()

	payload := &request.ChapterNovelRequest{
		Url: "https://kakuyomu.jp/works/16817330664532961874/episodes/16817330664868675561",
	}

	res, err := uts.kakuyomuService.KakuyomuChapterPage(payload)

	uts.Nil(res)
	uts.EqualError(err, "Bad Gateway Error")

	uts.MockTranslateRepository.AssertExpectations(uts.T())
}

func (uts *UnitTestSuite) TestKakuyomuChapterPage_ErrorColly() {
	payload := &request.ChapterNovelRequest{
		Url: "",
	}
	res, err := uts.kakuyomuService.KakuyomuChapterPage(payload)

	uts.Nil(res)
	uts.EqualError(err, "Failed to visit url")

	uts.MockTranslateRepository.AssertExpectations(uts.T())
}
