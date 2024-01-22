package service

import (
	"fmt"
	"jpnovelmtlgo/internal/exception"
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"jpnovelmtlgo/internal/repository"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/go-shiori/go-epub"
	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

type SyosetuServiceImpl struct {
	TranslateRepository repository.TranslateRepository
}

func NewSyosetuService(
	TranslateRepository repository.TranslateRepository,
) SyosetuService {
	return &SyosetuServiceImpl{
		TranslateRepository: TranslateRepository,
	}
}

func (service *SyosetuServiceImpl) ListChapterNovel(ctx *fiber.Ctx, params *request.ChapterNovelRequest) (*model.BaseResponse[[]request.TranslateListRequest], error) {
	var listChapter []request.TranslateListRequest

	c := colly.NewCollector()

	c.OnHTML(".index_box .novel_sublist2", func(e *colly.HTMLElement) {
		title := e.ChildText(".subtitle")
		url := e.ChildAttr("a", "href")
		urlSplit := strings.Split(url, "/")
		url = params.Url + urlSplit[2] + "/"
		chapter := &request.TranslateListRequest{
			Title: title,
			Url:   url,
		}

		listChapter = append(listChapter, *chapter)
	})

	err := c.Visit(params.Url)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	res, err := service.TranslateRepository.TranslateList(listChapter)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	sort.Slice(res.Data, func(i, j int) bool {
		return res.Data[i].Order < res.Data[j].Order
	})

	return res, nil
}

func (service *SyosetuServiceImpl) GetChapterPage(ctx *fiber.Ctx, params *request.ChapterNovelRequest) (*response.GetChapterPageResponse, error) {
	var title string
	var honbun string
	c := colly.NewCollector()

	c.OnHTML(".novel_subtitle", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML(".novel_view", func(e *colly.HTMLElement) {
		honbun = e.Text
	})

	err := c.Visit(params.Url)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	translateRequest := &request.TranslateChapterRequest{
		Title:   title,
		Chapter: honbun,
	}

	getTranslate, err := service.TranslateRepository.TranslateChapter(translateRequest)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	return getTranslate, nil
}

func (service *SyosetuServiceImpl) JpEpub(ctx *fiber.Ctx, params *request.ConvertNovelRequest) (*fiber.Map, error) {
	chapterNovel := make(chan request.ChapterContent)
	var listChapter []request.ChapterContent
	var wg sync.WaitGroup
	var title string
	var author string
	c := colly.NewCollector()

	c.OnHTML(".novel_title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML(".novel_writername", func(e *colly.HTMLElement) {
		author = e.Text
	})

	err := c.Visit(params.Url)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	e, err := epub.NewEpub(title)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}
	e.SetAuthor(author)
	pageSplit := strings.Split(params.Page, "-")
	startPage, _ := strconv.Atoi(pageSplit[0])
	finishPage, _ := strconv.Atoi(pageSplit[1])
	for i := startPage; i <= finishPage; i++ {
		wg.Add(1)
		payload := &request.ListChapterByUrl{
			Url:   params.Url + strconv.Itoa(i) + "/",
			Order: i,
		}
		go service.JpChapter(payload, &wg, chapterNovel)
	}

	go func() {
		wg.Wait()
		close(chapterNovel)
	}()

	for chapter := range chapterNovel {
		listChapter = append(listChapter, chapter)
	}

	sort.Slice(listChapter, func(i, j int) bool {
		return listChapter[i].Order < listChapter[j].Order
	})

	for _, item := range listChapter {
		sectionBody := `<h1>` + item.Title + `</h1>
		<p>` + item.Chapter + `</p>`
		_, err := e.AddSection(sectionBody, item.Title, "", "")
		if err != nil {
			panic(exception.GeneralError{
				Message: err.Error(),
			})
		}
	}

	err = e.Write(fmt.Sprintf("./epub/%s.epub", title))
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	res := &fiber.Map{
		"success": true,
	}
	return res, nil
}

func (service *SyosetuServiceImpl) JpChapter(params *request.ListChapterByUrl, wg *sync.WaitGroup, scrapingData chan<- request.ChapterContent) {
	defer wg.Done()
	var title string
	var honbun string
	c := colly.NewCollector()

	c.OnHTML(".novel_subtitle", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML(".novel_view", func(e *colly.HTMLElement) {
		honbun = e.Text
		honbun = strings.ReplaceAll(honbun, "\n", "<br />")
	})

	err := c.Visit(params.Url)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	content := request.ChapterContent{
		Title:   title,
		Chapter: honbun,
		Order:   params.Order,
	}

	scrapingData <- content
}

func (service *SyosetuServiceImpl) EnEpub(ctx *fiber.Ctx, params *request.ConvertNovelRequest) (*fiber.Map, error) {
	chapterNovel := make(chan request.ChapterContent)
	var listChapter []request.ChapterContent
	var wg sync.WaitGroup
	var title string
	var author string
	c := colly.NewCollector()

	c.OnHTML(".novel_title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML(".novel_writername", func(e *colly.HTMLElement) {
		author = e.Text
	})

	err := c.Visit(params.Url)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	novelInfo := &request.NovelInfo{
		Title:  title,
		Author: author,
	}
	translatedNovelInfo, err := service.TranslateRepository.TranslateInfo(novelInfo)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	e, err := epub.NewEpub(translatedNovelInfo.Title)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}
	e.SetAuthor(translatedNovelInfo.Author)
	pageSplit := strings.Split(params.Page, "-")
	startPage, _ := strconv.Atoi(pageSplit[0])
	finishPage, _ := strconv.Atoi(pageSplit[1])
	for i := startPage; i <= finishPage; i++ {
		wg.Add(1)
		payload := &request.ListChapterByUrl{
			Url:   params.Url + strconv.Itoa(i) + "/",
			Order: i,
		}
		go service.JpChapter(payload, &wg, chapterNovel)
	}

	go func() {
		wg.Wait()
		close(chapterNovel)
	}()

	for chapter := range chapterNovel {
		listChapter = append(listChapter, chapter)
	}

	sort.Slice(listChapter, func(i, j int) bool {
		return listChapter[i].Order < listChapter[j].Order
	})

	enBatch, err := service.TranslateRepository.TranslateListChapter(listChapter)
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	sort.Slice(enBatch, func(i, j int) bool {
		return enBatch[i].Order < enBatch[j].Order
	})

	for _, item := range enBatch {
		sectionBody := `<h1>` + item.Title + `</h1>
		<p>` + item.Chapter + `</p>`
		_, err := e.AddSection(sectionBody, item.Title, "", "")
		if err != nil {
			panic(exception.GeneralError{
				Message: err.Error(),
			})
		}
	}

	err = e.Write(fmt.Sprintf("./epub/%s.epub", translatedNovelInfo.Title))
	if err != nil {
		panic(exception.GeneralError{
			Message: err.Error(),
		})
	}

	res := &fiber.Map{
		"success": true,
	}
	return res, nil
}
