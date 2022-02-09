package errs

import "net/http"

type AppError struct {
	Code     int    `json:",omitempty"`
	Messsage string `json:"message"`
}

func (e *AppError) AsMessage() *AppError {
	return &AppError{
		Messsage: e.Messsage,
	}
}

func NewNotFoundError(msg string) *AppError {
	return &AppError{
		Code:     http.StatusNotFound,
		Messsage: msg,
	}
}

func NewInternalServerError(msg string) *AppError {
	return &AppError{
		Code:     http.StatusInternalServerError,
		Messsage: msg,
	}
}

func NewValidationError(msg string) *AppError {
	return &AppError{
		Code:     http.StatusUnprocessableEntity,
		Messsage: msg,
	}
}
