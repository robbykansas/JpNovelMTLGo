package service

import (
	"errors"
	"jpnovelmtlgo/internal/model"
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
	TranslateRepository repository.TranslateRepository,
) KakuyomuService {
	return &KakuyomuServiceImpl{
		TranslateRepository: TranslateRepository,
	}
}

func (service *KakuyomuServiceImpl) KakuyomuListChapter(params *request.ChapterNovelRequest) (*model.BaseResponse[[]request.TranslateListRequest], error) {
	var listChapter []request.TranslateListRequest

	c := colly.NewCollector()

	c.OnHTML("._workId__toc___I_tx ", func(e *colly.HTMLElement) {
		e.ForEach(".NewBox_box__45ont .WorkTocAccordion_contents__6nJhY", func(_ int, el *colly.HTMLElement) {
			el.ForEach(".NewBox_box__45ont", func(_ int, es *colly.HTMLElement) {
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
		return nil, errors.New("failed to visit url")
	}

	res, err := service.TranslateRepository.TranslateList(listChapter)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	sort.Slice(res.Data, func(i, j int) bool {
		return res.Data[i].Order < res.Data[j].Order
	})

	return res, nil
}

func (service *KakuyomuServiceImpl) KakuyomuChapterPage(params *request.ChapterNovelRequest) (*model.BaseResponse[*response.GetChapterPageResponse], error) {
	var title string
	var honbun string

	c := colly.NewCollector()

	c.OnHTML(".widget-episodeTitle", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML(".widget-episodeBody", func(e *colly.HTMLElement) {
		honbun = e.Text
	})

	err := c.Visit(params.Url)
	if err != nil {
		return nil, errors.New("Failed to visit url")
	}

	translateRequest := &request.TranslateChapterRequest{
		Title:   title,
		Chapter: honbun,
	}

	getTranslate, err := service.TranslateRepository.TranslateChapter(translateRequest)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	result := &model.BaseResponse[*response.GetChapterPageResponse]{
		StatusCode: "200",
		Message:    "Success",
		Data:       getTranslate,
	}

	return result, nil
}
