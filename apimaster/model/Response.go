package model

type ResponseSuccess struct {
	StatusCode int    `json:"statusCode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
