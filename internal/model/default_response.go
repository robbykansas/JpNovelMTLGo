package model

type DefaultResponse struct {
	IsSuccessful bool   `json:"isSuccessful"`
	StatusCode   string `json:"statusCode"`
	Message      string `json:"message"`
}

type BaseResponse[T any] struct {
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
}
