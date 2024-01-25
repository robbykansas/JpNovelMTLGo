package test

import (
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"

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

	uts.MockTranslateRepository.On("TranslateList", mock.AnythingOfType("[]request.TranslateListRequest")).Return(resultMock, nil)

	payload := &request.ChapterNovelRequest{
		Url: "https://ncode.syosetu.com/n6093en/",
	}

	res, err := uts.syosetuService.ListChapterNovel(payload)
	uts.Equal(resultMock, res)
	uts.Nil(err)
}

// func (uts *UnitTestSuite) TestListChapterNovel_Error() {
// 	mockData := []request.TranslateListRequest{
// 		{
// 			Title:   "example1",
// 			Url:     "url1",
// 			TitleEn: "exampleEn",
// 			Order:   1,
// 		},
// 		{
// 			Title:   "example2",
// 			Url:     "url2",
// 			TitleEn: "exampleEn2",
// 			Order:   2,
// 		},
// 	}

// 	resultMock := &model.BaseResponse[[]request.TranslateListRequest]{
// 		StatusCode: "200",
// 		Message:    "Success",
// 		Data:       mockData,
// 	}

// 	testError := errors.New("error")

// 	uts.MockTranslateRepository.On("TranslateList", mock.AnythingOfType("[]request.TranslateListRequest")).Return(nil, testError)

// 	payload := &request.ChapterNovelRequest{
// 		Url: "https://ncode.syosetu.com/n6093en/",
// 	}

// 	context := &fiber.Ctx{}

// 	res, err := uts.syosetuService.ListChapterNovel(context, payload)
// 	uts.Equal(resultMock, res)
// 	uts.EqualError(err, "Please wait a few minutes before you try again.")
// }

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
	uts.MockTranslateRepository.On("TranslateChapter", mock.AnythingOfType("*request.TranslateChapterRequest")).Return(mockTranslateResponse, nil)

	payload := &request.ChapterNovelRequest{
		Url: "http://ncode.syosetu.com/n6093en/395",
	}

	res, err := uts.syosetuService.GetChapterPage(payload)

	uts.Equal(mockResult, res)
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
