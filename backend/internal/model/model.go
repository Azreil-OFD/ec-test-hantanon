package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func SendJSONResponse(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		http.Error(w, "Ошибка при отправке ответа", http.StatusInternalServerError)
	}
}

const (
	DBError       = 0
	ConflictError = 1
	NotFound      = 2
)

type CustomError struct {
	Message string
	Code    int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %s (Code: %d)", e.Message, e.Code)
}
