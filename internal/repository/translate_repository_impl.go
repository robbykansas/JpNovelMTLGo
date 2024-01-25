package repository

import (
	"encoding/json"
	"fmt"
	"jpnovelmtlgo/internal/config"
	"jpnovelmtlgo/internal/exception"
	"jpnovelmtlgo/internal/model"
	"jpnovelmtlgo/internal/model/request"
	"jpnovelmtlgo/internal/model/response"
	"net/http"
	"strings"
	"sync"
	"time"
)

type TranslateRepositoryImpl struct {
	Configuration config.Config
}

func NewTranslateRepository(
	Configuration config.Config,
) TranslateRepository {
	return &TranslateRepositoryImpl{
		Configuration: Configuration,
	}
}

func (repository *TranslateRepositoryImpl) TranslateChapter(params *request.TranslateChapterRequest) (*response.GetChapterPageResponse, error) {
	translateTitle := make(chan string)
	translateChapter := make(chan string)
	errorChannel := make(chan error)

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

	go repository.TranslateWord(payloadTitleRequest, translateTitle, errorChannel)
	go repository.TranslateWord(payloadChapterRequest, translateChapter, errorChannel)

	select {
	case err := <-errorChannel:
		close(errorChannel)
		return nil, err
	default:
		title := <-translateTitle
		chapter := <-translateChapter
		result := &response.GetChapterPageResponse{
			Title:   title,
			Chapter: chapter,
		}

		close(translateTitle)
		close(translateChapter)

		return result, nil
	}
}

func (repository *TranslateRepositoryImpl) TranslateWord(params *request.TranslateRequest, channelWord chan<- string, errorChannel chan<- error) {
	client := &http.Client{}
	fmt.Println("<<<<<<<<<<<<<<<<<< access this")
	jsonData, err := json.Marshal(params)
	if err != nil {
		errorChannel <- err
	}

	payload := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", repository.Configuration.Get("TRANSLATE_URL"), payload)
	if err != nil {
		errorChannel <- err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		errorChannel <- err
	}

	translatedText := &response.TranslateResponse{}
	json.NewDecoder(res.Body).Decode(&translatedText)
	defer res.Body.Close()

	channelWord <- translatedText.TranslatedText
}

func (repository *TranslateRepositoryImpl) TranslateList(params []request.TranslateListRequest) (*model.BaseResponse[[]request.TranslateListRequest], error) {
	var wg sync.WaitGroup
	var translatedList []request.TranslateListRequest
	var count = 0
	translatedTitle := make(chan request.TranslateListRequest, 10)

	for _, item := range params {
		count += 1
		item.Order = count
		wg.Add(1)

		if count%50 == 0 {
			fmt.Println(count)
			time.Sleep(10 * time.Second)
		}

		go repository.TranslateEachTitle(item, &wg, translatedTitle)
	}

	go func() {
		wg.Wait()
		close(translatedTitle)
	}()

	for title := range translatedTitle {
		translatedList = append(translatedList, title)
	}

	result := &model.BaseResponse[[]request.TranslateListRequest]{
		StatusCode: "200",
		Message:    "Success",
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

func (repository *TranslateRepositoryImpl) TranslateInfo(params *request.NovelInfo) (*response.TranslatedInfoResponse, error) {
	translatedTitle := make(chan string)
	translatedAuthor := make(chan string)
	errorChannel := make(chan error)

	payloadTitle := &request.TranslateRequest{
		Q:      params.Title,
		Source: "ja",
		Target: "en",
		Format: "",
	}

	payloadAuthor := request.TranslateRequest{
		Q:      params.Author,
		Source: "ja",
		Target: "en",
		Format: "",
	}

	go repository.TranslateWord(payloadTitle, translatedTitle, errorChannel)
	go repository.TranslateWord(&payloadAuthor, translatedAuthor, errorChannel)

	select {
	case err := <-errorChannel:
		close(errorChannel)
		return nil, err
	default:
		title := <-translatedTitle
		author := <-translatedAuthor
		result := &response.TranslatedInfoResponse{
			Title:  title,
			Author: author,
		}

		close(translatedTitle)
		close(translatedAuthor)

		return result, nil
	}
}

func (repository *TranslateRepositoryImpl) TranslateListChapter(params []request.ChapterContent) ([]response.TranslatedListChapterResponse, error) {
	var wg sync.WaitGroup
	var listChapter []response.TranslatedListChapterResponse
	translatedChapter := make(chan response.TranslatedListChapterResponse, 10)

	for _, item := range params {
		wg.Add(1)

		go repository.TranslateEachChapter(item, &wg, translatedChapter)
	}

	go func() {
		wg.Wait()
		close(translatedChapter)
	}()

	for chapter := range translatedChapter {
		listChapter = append(listChapter, chapter)
	}

	return listChapter, nil
}

func (repository *TranslateRepositoryImpl) TranslateEachChapter(params request.ChapterContent, wg *sync.WaitGroup, translatedChapter chan<- response.TranslatedListChapterResponse) {
	defer wg.Done()

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

	enTitle, err := repository.SyncTranslateWord(payloadTitleRequest)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	enChapter, err := repository.SyncTranslateWord(payloadChapterRequest)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	translatedChapter <- response.TranslatedListChapterResponse{
		Title:   enTitle,
		Chapter: enChapter,
		Order:   params.Order,
	}
}

func (repository *TranslateRepositoryImpl) SyncTranslateWord(params *request.TranslateRequest) (string, error) {
	client := &http.Client{}

	jsonData, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	payload := strings.NewReader(string(jsonData))

	req, err := http.NewRequest("POST", repository.Configuration.Get("TRANSLATE_URL"), payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	translatedText := &response.TranslateResponse{}
	json.NewDecoder(res.Body).Decode(&translatedText)
	defer res.Body.Close()

	return translatedText.TranslatedText, nil
}
