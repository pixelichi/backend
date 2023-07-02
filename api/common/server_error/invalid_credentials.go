package server_error

import "net/http"

type InvalidCredentialsError struct {
	BaseErrorResponse
}


func NewInvalidCredentialsError(detail string) CustomError {
	if detail == "" {
		detail = "Invalid Credentials, please check your username and password."
	}

	return &InvalidCredentialsError{
		BaseErrorResponse{
			Type:     "https://shinypothos.com/docs/errors/invalid-credentials",
			Title:    "Invalid Credentials",
			Status:   http.StatusNotFound,
			Detail:   detail,
			Instance: "", // You can assign an instance value or leave it empty, based on your requirements.
		},
	}
}