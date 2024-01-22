// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"

	model "jpnovelmtlgo/internal/model"

	request "jpnovelmtlgo/internal/model/request"

	response "jpnovelmtlgo/internal/model/response"
)

// KakuyomuService is an autogenerated mock type for the KakuyomuService type
type KakuyomuService struct {
	mock.Mock
}

// KakuyomuChapterPage provides a mock function with given fields: ctx, params
func (_m *KakuyomuService) KakuyomuChapterPage(ctx *fiber.Ctx, params *request.ChapterNovelRequest) (*response.GetChapterPageResponse, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for KakuyomuChapterPage")
	}

	var r0 *response.GetChapterPageResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, *request.ChapterNovelRequest) (*response.GetChapterPageResponse, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, *request.ChapterNovelRequest) *response.GetChapterPageResponse); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*response.GetChapterPageResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*fiber.Ctx, *request.ChapterNovelRequest) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KakuyomuListChapter provides a mock function with given fields: ctx, params
func (_m *KakuyomuService) KakuyomuListChapter(ctx *fiber.Ctx, params *request.ChapterNovelRequest) (*model.BaseResponse[[]request.TranslateListRequest], error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for KakuyomuListChapter")
	}

	var r0 *model.BaseResponse[[]request.TranslateListRequest]
	var r1 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, *request.ChapterNovelRequest) (*model.BaseResponse[[]request.TranslateListRequest], error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, *request.ChapterNovelRequest) *model.BaseResponse[[]request.TranslateListRequest]); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BaseResponse[[]request.TranslateListRequest])
		}
	}

	if rf, ok := ret.Get(1).(func(*fiber.Ctx, *request.ChapterNovelRequest) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewKakuyomuService creates a new instance of KakuyomuService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKakuyomuService(t interface {
	mock.TestingT
	Cleanup(func())
}) *KakuyomuService {
	mock := &KakuyomuService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
