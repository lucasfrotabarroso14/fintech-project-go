package responses

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(w http.ResponseWriter, status int, Message string, Data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&JSONResponse{
		Status:  status,
		Message: Message,
		Data:    Data,
	})
}

func Error(w http.ResponseWriter, status int, message string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&JSONResponse{
		Status:  status,
		Message: message,
		Error:   err.Error(),
	})
}
