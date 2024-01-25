package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
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

func (uts *UnitTestSuite) TestKakuyomuListChapterController_error() {
	errData := errors.New("error data")

	payload := &request.ChapterNovelRequest{
		Url: "https://kakuyomu.jp/works/16817139558533391541",
	}

	uts.MockKakuyomuService.On("KakuyomuListChapter", payload).Return(nil, errData)

	assert.Panics(uts.T(), func() {
		uts.kakuyomuController.KakuyomuListChapter(&fiber.Ctx{})
	})
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

	fmt.Println(resp, "<<<< resp")
	uts.Equal(http.StatusOK, resp.StatusCode)
}

// func (uts *UnitTestSuite) TestKakuyomuChapterPageController_error() {
// 	errData := errors.New("error data")

// 	payload := &request.ChapterNovelRequest{
// 		Url: "https://kakuyomu.jp/works/16817139558533391541",
// 	}
// 	uts.MockKakuyomuService.On("KakuyomuChapterPage", payload).Return(nil, errData)

// 	payloadByte, _ := json.Marshal(payload)
// 	body := string(payloadByte)
// 	req := httptest.NewRequest(http.MethodPost, "/kakuyomu/chapter", strings.NewReader(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	_, err := uts.app.Test(req)
// 	if err != nil {
// 		uts.Fail("Expected to get error", err)
// 	}
// }
