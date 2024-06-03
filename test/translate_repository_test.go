package test

import (
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
)

func (uts *UnitTestSuite) TestTranslateChapter() {
	payload := &request.TranslateChapterRequest{
		Title:   "1 メンバー募集",
		Chapter: "5 帝都",
	}

	resultTest := &response.GetChapterPageResponse{
		Title:   "1 Recruitment",
		Chapter: "5 Chengdu",
	}

	uts.MockConfig.On("Get", "TRANSLATE_URL").Return("http://127.0.0.1:5001/translate")

	res, err := uts.translateRepository.TranslateChapter(payload)
	uts.Equal(resultTest, res)
	uts.Nil(err)
}
