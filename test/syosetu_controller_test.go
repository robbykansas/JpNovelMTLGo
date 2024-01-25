package test

import (
	"encoding/json"
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (uts *UnitTestSuite) TestListChapterNovelController() {
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
		Url: "https://ncode.syosetu.com/n6093en",
	}
	uts.MockSyosetuService.On("ListChapterNovel", payload).Return(resultMock, nil)

	payloadByte, _ := json.Marshal(payload)
	body := string(payloadByte)
	req := httptest.NewRequest(http.MethodPost, "/syosetu/list-chapter", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := uts.app.Test(req)

	uts.Equal(http.StatusOK, resp.StatusCode)
}

func (uts *UnitTestSuite) TestGetChapterController() {
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
		Url: "https://ncode.syosetu.com/n6093en/395",
	}
	uts.MockSyosetuService.On("GetChapterPage", payload).Return(mockResult, nil)

	payloadByte, _ := json.Marshal(payload)
	body := string(payloadByte)
	req := httptest.NewRequest(http.MethodPost, "/syosetu/chapter", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := uts.app.Test(req)

	uts.Equal(http.StatusOK, resp.StatusCode)
}

func (uts *UnitTestSuite) TestJpEpubController() {
	mockResult := &model.DefaultResponse{
		StatusCode:   "200",
		Message:      "Success",
		IsSuccessful: true,
	}

	payload := &request.ConvertNovelRequest{
		Url:  "https://ncode.syosetu.com/n6093en",
		Page: "1-50",
	}
	uts.MockSyosetuService.On("JpEpub", payload).Return(mockResult, nil)

	payloadByte, _ := json.Marshal(payload)
	body := string(payloadByte)
	req := httptest.NewRequest(http.MethodPost, "/syosetu/jp-epub", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := uts.app.Test(req)

	uts.Equal(http.StatusOK, resp.StatusCode)
}

func (uts *UnitTestSuite) TestEnEpubController() {
	mockResult := &model.DefaultResponse{
		StatusCode:   "200",
		Message:      "Success",
		IsSuccessful: true,
	}

	payload := &request.ConvertNovelRequest{
		Url:  "https://ncode.syosetu.com/n6093en",
		Page: "1-50",
	}
	uts.MockSyosetuService.On("EnEpub", payload).Return(mockResult, nil)

	payloadByte, _ := json.Marshal(payload)
	body := string(payloadByte)
	req := httptest.NewRequest(http.MethodPost, "/syosetu/en-epub", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := uts.app.Test(req)

	uts.Equal(http.StatusOK, resp.StatusCode)
}
