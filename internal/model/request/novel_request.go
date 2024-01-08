package request

type ListChapterNovelRequest struct {
	Url string `json:"url"`
}

type TranslateChapterRequest struct {
	Title   string `json:"title"`
	Chapter string `json:"chapter"`
}
