package errors

import (
	"fmt"
	"net/http"
)

type ApiError interface {
	Message() string
	Code() string
	Status() int
	Error() string
}

type apiErr struct {
	ErrorMessage string    `json:"message"`
	ErrorCode    string    `json:"errors"`
	ErrorStatus  int       `json:"status"`
}

func NewApiError(message string, error string, status int) ApiError {
	return apiErr{message, error, status}
}

func NewNotFoundApiError(message string) ApiError {
	return apiErr{message, "not_found", http.StatusNotFound}
}

func NewInternalServerApiError(message string, err error) ApiError {
	return apiErr{message, "internal_server_error", http.StatusInternalServerError}
}

func NewBadRequestApiError(message string) ApiError {
	return apiErr{message, "bad_request", http.StatusBadRequest}
}

func (e apiErr) Code() string {
	return e.ErrorCode
}

func (e apiErr) Error() string {
	return fmt.Sprintf("Message: %s;Error Code: %s;Status: %d", e.ErrorMessage, e.ErrorCode, e.ErrorStatus)
}

func (e apiErr) Status() int {
	return e.ErrorStatus
}

func (e apiErr) Message() string {
	return e.ErrorMessage
}