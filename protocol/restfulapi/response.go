package restfulapi

import (
	"encoding/json"
	"go-live/models"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type AppsResponse struct {
	Code    int          `json:"code"`
	Data    []models.App `json:"data"`
	Message string       `json:"message"`
}

func SendErrorResponse(w http.ResponseWriter, code int, message string) {
	SendResponse(w, &ErrorResponse{
		Code:    code,
		Message: message,
	})
}

func SendResponse(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)

	w.Write(data)
}
