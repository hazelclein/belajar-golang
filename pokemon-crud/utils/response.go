package utils

import (
	"encoding/json"
	"net/http"
)

// struct response sukses
type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// struct response error
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// func kirim response sukses
func SendSuccess(w http.ResponseWriter, message string, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	response := SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	
	json.NewEncoder(w).Encode(response)
}

// func kirim response error
func SendError(w http.ResponseWriter, message string, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}
	
	response := ErrorResponse{
		Status:  "error",
		Message: message,
		Error:   errorMsg,
	}
	
	json.NewEncoder(w).Encode(response)
}