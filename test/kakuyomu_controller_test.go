package test

import (
	"encoding/json"
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (uts *UnitTestSuite) TestKakuyomuListChapterController() {
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

	payload := &request.ChapterNovelRequest{
		Url: "https://kakuyomu.jp/works/16817139558533391541",
	}
	uts.MockKakuyomuService.On("KakuyomuListChapter", payload).Return(resultMock, nil)

	payloadByte, _ := json.Marshal(payload)
	body := string(payloadByte)
	req := httptest.NewRequest(http.MethodPost, "/kakuyomu/list-chapter", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := uts.app.Test(req)

	uts.Equal(http.StatusOK, resp.StatusCode)
}

func (uts *UnitTestSuite) TestKakuyomuChapterPageController() {
	mockTranslateResponse := &response.GetChapterPageResponse{
		Title:   "titleEn",
		Chapter: "chapterEn",
	}
	mockResult := &model.BaseResponse[*response.GetChapterPageResponse]{
		StatusCode: "200",
		Message:    "Success",
		Data:       mockTranslateResponse,
	}

	payload := &request.ChapterNovelRequest{
		Url: "https://kakuyomu.jp/works/16817139558533391541",
	}
	uts.MockKakuyomuService.On("KakuyomuChapterPage", payload).Return(mockResult, nil)

	payloadByte, _ := json.Marshal(payload)
	body := string(payloadByte)
	req := httptest.NewRequest(http.MethodPost, "/kakuyomu/chapter", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := uts.app.Test(req)
	uts.Equal(http.StatusOK, resp.StatusCode)
}

func (uts *UnitTestSuite) TestKakuyomuChapterPage_Error() {
	errData := fiber.NewError(fiber.StatusBadRequest, "error mock")

	payload := &request.ChapterNovelRequest{
		Url: "https://kakuyomu.jp/works/16817330664532961874",
	}
	uts.MockKakuyomuService.On("KakuyomuChapterPage", payload).Return(nil, errData)

	payloadByte, _ := json.Marshal(payload)
	body := string(payloadByte)
	req := httptest.NewRequest(http.MethodPost, "/kakuyomu/chapter", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := uts.app.Test(req)
	uts.Equal(http.StatusBadRequest, resp.StatusCode)
}
