package request

type ListChapterNovelRequest struct {
	Url string `json:"url"`
}

type TranslateChapterRequest struct {
	Title   string `json:"title"`
	Chapter string `json:"chapter"`
}

type TranslateListRequest struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	TitleEn string `json:"titleEn"`
	Order   int    `json:"order"`
}
