package response

type ListChapterNovel struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type ListChapterNovelResponse struct {
	StatusCode string             `json:"statusCode"`
	Data       []ListChapterNovel `json:"data"`
}

type GetChapterPageResponse struct {
	Title   string `json:"title"`
	Chapter string `json:"chapter"`
}
