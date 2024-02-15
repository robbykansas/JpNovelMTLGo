package test

import (
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

func (uts *UnitTestSuite) TestListChapterNovel() {
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

	resultMock := &model.BaseResponse[[]request.TranslateListRequest]{
		StatusCode: "200",
		Message:    "Success",
		Data:       mockData,
	}

	uts.MockTranslateRepository.On(
		"TranslateList",
		mock.AnythingOfType("[]request.TranslateListRequest")).
		Return(resultMock, nil).Once()

	payload := &request.ChapterNovelRequest{
		Url: "https://ncode.syosetu.com/n6093en/",
	}

	res, err := uts.syosetuService.ListChapterNovel(payload)
	uts.Equal(resultMock, res)
	uts.Nil(err)

	uts.MockTranslateRepository.ExpectedCalls = nil

	errData := fiber.NewError(fiber.StatusBadGateway, "Bad Gateway Error")
	uts.MockTranslateRepository.On(
		"TranslateList",
		mock.AnythingOfType("[]request.TranslateListRequest")).
		Return(nil, errData).Once()

	resNil, err := uts.syosetuService.ListChapterNovel(payload)

	uts.Nil(resNil)
	uts.EqualError(err, "Bad Gateway Error")
}

func (uts *UnitTestSuite) TestGetChapterPage() {
	mockTranslateResponse := &response.GetChapterPageResponse{
		Title:   "titleEn",
		Chapter: "chapterEn",
	}

	mockResult := &model.BaseResponse[*response.GetChapterPageResponse]{
		StatusCode: "200",
		Message:    "Success",
		Data:       mockTranslateResponse,
	}
	uts.MockTranslateRepository.On(
		"TranslateChapter",
		mock.AnythingOfType("*request.TranslateChapterRequest")).
		Return(mockTranslateResponse, nil).Once()

	payload := &request.ChapterNovelRequest{
		Url: "http://ncode.syosetu.com/n6093en/395",
	}

	res, err := uts.syosetuService.GetChapterPage(payload)

	uts.Equal(mockResult, res)
	uts.Nil(err)

	uts.MockTranslateRepository.ExpectedCalls = nil

	errData := fiber.NewError(fiber.StatusBadGateway, "Bad Gateway Error")
	uts.MockTranslateRepository.On(
		"TranslateChapter",
		mock.AnythingOfType("*request.TranslateChapterRequest")).
		Return(nil, errData).Once()

	resNil, err := uts.syosetuService.GetChapterPage(payload)

	uts.Nil(resNil)
	uts.EqualError(err, "Bad Gateway Error")
}

func (uts *UnitTestSuite) TestListChapterNovel_ErrorURL() {
	payload := &request.ChapterNovelRequest{
		Url: "",
	}

	res, err := uts.syosetuService.ListChapterNovel(payload)
	uts.Nil(res)
	uts.EqualError(err, "Failed to visit url")
}

func (uts *UnitTestSuite) TestGetChapterPage_ErrorURL() {
	payload := &request.ChapterNovelRequest{
		Url: "",
	}

	res, err := uts.syosetuService.GetChapterPage(payload)
	uts.Nil(res)
	uts.EqualError(err, "Failed to visit url")
}
