package server_error

import "net/http"

type NotFoundError struct {
	BaseErrorResponse
}


func NewNotFoundError(detail string) CustomError {
	if detail == "" {
		detail = "Not found"
	}

	return &InvalidCredentialsError{
		BaseErrorResponse{
			Type:     "https://shinypothos.com/docs/errors/invalid-credentials",
			Title:    "Not found",
			Status:   http.StatusNotFound,
			Detail:   detail,
			Instance: "", // You can assign an instance value or leave it empty, based on your requirements.
		},
	}
}