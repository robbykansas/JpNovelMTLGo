package response

import "jpnovelmtlgo/internal/model/request"

type ListChapterNovelResponse struct {
	StatusCode string                         `json:"statusCode"`
	Data       []request.TranslateListRequest `json:"data"`
}

type EnBatchChapterResponse struct {
	StatusCode string                   `json:"statusCode"`
	Data       []request.ChapterContent `json:"data"`
}

type GetChapterPageResponse struct {
	Title   string `json:"title"`
	Chapter string `json:"chapter"`
}

type TranslatedInfoResponse struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type TranslatedListChapterResponse struct {
	Title   string `json:"title"`
	Chapter string `json:"chapter"`
	Order   int    `json:"order"`
}
