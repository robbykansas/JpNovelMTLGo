package request

type ChapterNovelRequest struct {
	Url string `json:"url"`
}

type ConvertNovelRequest struct {
	Url  string `json:"url"`
	Page string `json:"page"`
}

type ListChapterByUrl struct {
	Url   string `json:"url"`
	Order int    `json:"order"`
}

type ChapterContent struct {
	Title   string `json:"title"`
	Chapter string `json:"chapter"`
	Order   int    `json:"order"`
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
