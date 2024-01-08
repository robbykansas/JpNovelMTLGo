package repository

import (
	"encoding/json"
	"jpnovelmtlgo/internal/config"
	"jpnovelmtlgo/internal/exception"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"net/http"
	"strings"
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

		req, err := http.NewRequest("POST", repository.Configuration.App().Translate.Url, payload)
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

		req, err := http.NewRequest("POST", repository.Configuration.App().Translate.Url, payload)
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

func (repository *TranslateRepositoryImpl) AsyncTranslate(params *request.TranslateRequest) string {
	client := &http.Client{}

	jsonData, err := json.Marshal(params)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	payload := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", repository.Configuration.App().Translate.Url, payload)
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

	return translatedText.TranslatedText
}
