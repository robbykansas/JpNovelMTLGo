package service

import (
	"jpnovelmtlgo/internal/exception"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"jpnovelmtlgo/internal/repository"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

type SyosetuServiceImpl struct {
	TranslateRepository repository.TranslateRepository
}

func NewSyosetuService(
	TranslateRepository *repository.TranslateRepository,
) SyosetuService {
	return &SyosetuServiceImpl{
		TranslateRepository: *TranslateRepository,
	}
}

func (service *SyosetuServiceImpl) ListChapterNovel(ctx *fiber.Ctx, params *request.ListChapterNovelRequest) (*response.ListChapterNovelResponse, error) {
	var listChapter []response.ListChapterNovel

	c := colly.NewCollector()

	c.OnHTML(".index_box .novel_sublist2", func(e *colly.HTMLElement) {
		title := e.ChildText(".subtitle")
		url := e.ChildAttr("a", "href")
		urlSplit := strings.Split(url, "/")
		url = params.Url + urlSplit[2] + "/"
		chapter := &response.ListChapterNovel{
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
	response := &response.ListChapterNovelResponse{
		StatusCode: "200",
		Data:       listChapter,
	}

	return response, nil
}

func (service *SyosetuServiceImpl) GetChapterPage(ctx *fiber.Ctx, params *request.ListChapterNovelRequest) (*response.GetChapterPageResponse, error) {
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
