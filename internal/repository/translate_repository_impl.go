package repository

import (
	"encoding/json"
	"jpnovelmtlgo/internal/config"
	"jpnovelmtlgo/internal/exception"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"net/http"
	"strings"
	"sync"
)

type TranslateRepositoryImpl struct {
	Configuration config.Config
}

func NewTranslateRepository(
	Configuration *config.Config,
) TranslateRepository {
	return &TranslateRepositoryImpl{
		Configuration: *Configuration,
	}
}

func (repository *TranslateRepositoryImpl) TranslateChapter(params *request.TranslateChapterRequest) (*response.GetChapterPageResponse, error) {
	translateTitle := make(chan string)
	translateChapter := make(chan string)

	payloadTitleRequest := &request.TranslateRequest{
		Q:      params.Title,
		Source: "ja",
		Target: "en",
		Format: "",
	}

	payloadChapterRequest := &request.TranslateRequest{
		Q:      params.Chapter,
		Source: "ja",
		Target: "en",
		Format: "",
	}

	go func() {
		client := &http.Client{}

		jsonData, err := json.Marshal(payloadTitleRequest)
		if err != nil {
			exception.PanicIfNeeded(err)
		}

		payload := strings.NewReader(string(jsonData))

		req, err := http.NewRequest("POST", repository.Configuration.Get("TRANSLATE_URL"), payload)
		if err != nil {
			exception.PanicIfNeeded(err)
		}

		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			exception.PanicIfNeeded(err)
		}

		translatedText := &response.TranslateResponse{}
		json.NewDecoder(res.Body).Decode(&translatedText)
		defer res.Body.Close()

		translateTitle <- translatedText.TranslatedText
		defer close(translateTitle)
	}()

	go func() {
		client := &http.Client{}

		jsonData, err := json.Marshal(payloadChapterRequest)
		if err != nil {
			exception.PanicIfNeeded(err)
		}

		payload := strings.NewReader(string(jsonData))

		req, err := http.NewRequest("POST", repository.Configuration.Get("TRANSLATE_URL"), payload)
		if err != nil {
			exception.PanicIfNeeded(err)
		}

		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			exception.PanicIfNeeded(err)
		}

		translatedText := &response.TranslateResponse{}
		json.NewDecoder(res.Body).Decode(&translatedText)
		defer res.Body.Close()

		translateChapter <- translatedText.TranslatedText
		defer close(translateChapter)
	}()

	title := <-translateTitle
	chapter := <-translateChapter
	result := &response.GetChapterPageResponse{
		Title:   title,
		Chapter: chapter,
	}

	return result, nil
}

func (repository *TranslateRepositoryImpl) TranslateList(params []request.TranslateListRequest) (*response.ListChapterNovelResponse, error) {
	var wg sync.WaitGroup
	var translatedList []request.TranslateListRequest
	var count = 0
	translatedTitle := make(chan request.TranslateListRequest, 10)

	for _, item := range params {
		count += 1
		item.Order = count
		wg.Add(1)

		go repository.TranslateEachTitle(item, &wg, translatedTitle)
	}

	wg.Wait()
	close(translatedTitle)

	for title := range translatedTitle {
		translatedList = append(translatedList, title)
	}

	result := &response.ListChapterNovelResponse{
		StatusCode: "200",
		Data:       translatedList,
	}

	return result, nil
}

func (repository *TranslateRepositoryImpl) TranslateEachTitle(params request.TranslateListRequest, wg *sync.WaitGroup, translatedTitle chan<- request.TranslateListRequest) {
	defer wg.Done()
	client := &http.Client{}

	payloadTitleRequest := &request.TranslateRequest{
		Q:      params.Title,
		Source: "ja",
		Target: "en",
		Format: "",
	}

	jsonData, err := json.Marshal(payloadTitleRequest)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	payload := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", repository.Configuration.Get("TRANSLATE_URL"), payload)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	translatedText := &response.TranslateResponse{}
	json.NewDecoder(res.Body).Decode(&translatedText)
	defer res.Body.Close()

	translatedTitle <- request.TranslateListRequest{
		Title:   params.Title,
		Url:     params.Url,
		TitleEn: translatedText.TranslatedText,
		Order:   params.Order,
	}
}
