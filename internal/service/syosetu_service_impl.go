package service

import (
	"jpnovelmtlgo/internal/exception"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

type SyosetuServiceImpl struct{}

func NewSyosetuService() SyosetuService {
	return &SyosetuServiceImpl{}
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
