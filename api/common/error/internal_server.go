package error

import "net/http"

type InternalServerError struct {
	BaseErrorResponse
}

func NewInternalServerError(detail string) CustomError {
	if detail == "" {
		detail = "Internal Server Error, something went wrong."
	}

	return &InvalidCredentialsError{
		BaseErrorResponse{
			Type:     "https://pixelichi.com/docs/errors/internal-server-error",
			Title:    "Internal Server Error",
			Status:   http.StatusInternalServerError,
			Detail:   detail,
			Instance: "", // You can assign an instance value or leave it empty, based on your requirements.
		},
	}
}
