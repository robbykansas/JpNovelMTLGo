package response

import "jpnovelmtlgo/internal/model/request"

type ListChapterNovelResponse struct {
	StatusCode string                         `json:"statusCode"`
	Data       []request.TranslateListRequest `json:"data"`
}

type GetChapterPageResponse struct {
	Title   string `json:"title"`
	Chapter string `json:"chapter"`
}
