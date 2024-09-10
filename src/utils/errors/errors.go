package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type ApiError struct {
	Status  int
	Message string
}

func GetStatus(err error) int {
	status := http.StatusInternalServerError

	var apiErr ApiError
	if errors.As(err, &apiErr) {
		status = apiErr.Status
	}
	return status
}

func GetMessage(err error) string {
	msg := err.Error()

	var apiErr ApiError
	if errors.As(err, &apiErr) {
		msg = apiErr.Message
	}
	return msg
}

func NewApiError(status int, message string) error {
	return ApiError{
		Status:  status,
		Message: message,
	}
}

func (a ApiError) Error() string {
	return fmt.Sprintf("%s: %s", http.StatusText(a.Status), a.Message)
}
