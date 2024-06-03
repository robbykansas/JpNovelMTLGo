package request

type ChapterNovelRequest struct {
	Url string `json:"url"`
}

type ChapterNovelListPageRequest struct {
	Url  string `json:"url"`
	Page int    `json:"page"`
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
	Page    int    `json:"page"`
}

type TranslateListEach struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	TitleEn string `json:"titleEn"`
	Order   int    `json:"order"`
	Page    int    `json:"page"`
}

type NovelInfo struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
