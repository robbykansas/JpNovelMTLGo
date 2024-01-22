package response

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
