package server_error

import "net/http"

type BadRequestError struct {
	BaseErrorResponse
}


func NewBadRequestError(detail string) CustomError {
	if detail == "" {
		detail = "Bad request"
	}

	return &BadRequestError{
		BaseErrorResponse{
			Type:     "https://shinypothos.com/docs/errors/bad-request",
			Title:    "Bad Request",
			Status:   http.StatusUnauthorized,
			Detail:   detail,
			Instance: "", // You can assign an instance value or leave it empty, based on your requirements.
		},
	}
}