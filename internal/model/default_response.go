package model

type DefaultResponse struct {
	IsSuccessful bool   `json:"isSuccessful"`
	StatusCode   string `json:"statusCode"`
	Message      string `json:"message"`
}
