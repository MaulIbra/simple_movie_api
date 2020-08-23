package util

import (
	"encoding/json"
	"github.com/maulibra/simple_movie_api/apimaster/model"
	"net/http"
)

const (
	RESPONSE_SUCCESS  = "Success"
	BAD_REQUEST       = "Bad Request"
	BAD_GATEWAY       = "Bad Gateway"
	NOT_FOUND         = "Not Found"
	STATUS_NO_CONTENT = "No Content"
	STATUS_CREATED    = "Created"
)

func messageStatusCode(statusCode int) string {
	switch statusCode {
	case http.StatusBadRequest:
		return BAD_REQUEST
	case http.StatusBadGateway:
		return BAD_GATEWAY
	case http.StatusNotFound:
		return NOT_FOUND
	case http.StatusOK:
		return RESPONSE_SUCCESS
	case http.StatusCreated:
		return STATUS_CREATED
	default:
		return BAD_REQUEST
	}
}


func HandleResponseError(res http.ResponseWriter, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response := model.ResponseError{StatusCode: statusCode, Message: http.StatusText(statusCode)}
	responseObj, _ := json.Marshal(response)
	res.Write(responseObj)
}

func HandleResponseSuccess(res http.ResponseWriter, statusCode int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	response := model.ResponseSuccess{StatusCode: statusCode, Message: http.StatusText(statusCode), Data: data}
	byteOfData, _ := json.Marshal(response)
	res.Write(byteOfData)
}
