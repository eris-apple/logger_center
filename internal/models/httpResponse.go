package models

type HttpResponseError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type HttpResponseSuccess struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}
