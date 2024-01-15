package service

import (
	"jpnovelmtlgo/internal/exception"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"jpnovelmtlgo/internal/repository"
	"sort"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
)

type KakuyomuServiceImpl struct {
	TranslateRepository repository.TranslateRepository
}

func NewKakuyomuService(
	TranslateRepository *repository.TranslateRepository,
) KakuyomuService {
	return &KakuyomuServiceImpl{
		TranslateRepository: *TranslateRepository,
	}
}

func (service *KakuyomuServiceImpl) KakuyomuListChapter(ctx *fiber.Ctx, params *request.ChapterNovelRequest) (*response.ListChapterNovelResponse, error) {
	var listChapter []request.TranslateListRequest

	c := colly.NewCollector()

	c.OnHTML("._workId__toc___I_tx ", func(e *colly.HTMLElement) {
		// default1 := ".WorkTocSection_link__ocg9K"
		// default2 := ".WorkTocSection_title__H2007"
		e.ForEach(".NewBox_box__45ont .WorkTocAccordion_contents__6nJhY", func(_ int, el *colly.HTMLElement) {
			el.ForEach(".WorkTocSection_link__ocg9K", func(_ int, es *colly.HTMLElement) {
				title := es.ChildText(".WorkTocSection_title__H2007")
				url := es.Attr("href")

				chapter := &request.TranslateListRequest{
					Title: title,
					Url:   url,
				}

				listChapter = append(listChapter, *chapter)
			})
		})
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
