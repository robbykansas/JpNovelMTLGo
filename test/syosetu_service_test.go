package test

import (
	"errors"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"jpnovelmtlgo/internal/service"
	"jpnovelmtlgo/mocks"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UnitTestSuite struct {
	suite.Suite
	syosetuService          service.SyosetuService
	kakuyomuService         service.KakuyomuService
	MockTranslateRepository *mocks.TranslateRepository
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (uts *UnitTestSuite) SetupSuite() {
	MockTranslateRepository := mocks.TranslateRepository{}
	syosetuService := service.NewSyosetuService(&MockTranslateRepository)
	kakuyomuService := service.NewKakuyomuService(&MockTranslateRepository)

	uts.syosetuService = syosetuService
	uts.kakuyomuService = kakuyomuService
	uts.MockTranslateRepository = &MockTranslateRepository
}

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

	resultMock := &response.ListChapterNovelResponse{
		StatusCode: "200",
		Data:       mockData,
	}

	uts.MockTranslateRepository.On("TranslateList", mock.AnythingOfType("[]request.TranslateListRequest")).Return(resultMock, nil)

	payload := &request.ChapterNovelRequest{
		Url: "https://ncode.syosetu.com/n6093en/",
	}

	context := &fiber.Ctx{}

	res, err := uts.syosetuService.ListChapterNovel(context, payload)
	uts.Equal(resultMock, res)
	uts.Nil(err)
}

func (uts *UnitTestSuite) TestListChapterNovel_Error() {
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

	testError := errors.New("error")

	uts.MockTranslateRepository.On("TranslateList", mock.AnythingOfType("[]request.TranslateListRequest")).Return(nil, testError)

	payload := &request.ChapterNovelRequest{
		Url: "https://ncode.syosetu.com/n6093en/",
	}

	context := &fiber.Ctx{}

	res, err := uts.syosetuService.ListChapterNovel(context, payload)
	uts.Equal(resultMock, res)
	uts.EqualError(err, "Please wait a few minutes before you try again.")
}

func (uts *UnitTestSuite) TestGetChapterPage() {
	mockTranslateResponse := &response.GetChapterPageResponse{
		Title:   "titleEn",
		Chapter: "chapterEn",
	}
	uts.MockTranslateRepository.On("TranslateChapter", mock.AnythingOfType("*request.TranslateChapterRequest")).Return(mockTranslateResponse, nil)

	payload := &request.ChapterNovelRequest{
		Url: "http://ncode.syosetu.com/n6093en/395",
	}
	context := &fiber.Ctx{}

	res, err := uts.syosetuService.GetChapterPage(context, payload)

	uts.Equal(mockTranslateResponse, res)
	uts.Nil(err)
}

// func (uts *UnitTestSuite) TestJpEpub() {
// 	payload := &request.ConvertNovelRequest{
// 		Url:  "http://ncode.syosetu.com/n6093en/",
// 		Page: "1-2",
// 	}
// 	context := &fiber.Ctx{}

// 	result := &fiber.Map{
// 		"success": true,
// 	}

// 	res, err := uts.syosetuService.JpEpub(context, payload)
// 	uts.Equal(result, res)
// 	uts.Nil(err)
// }
