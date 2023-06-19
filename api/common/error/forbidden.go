package error

import "net/http"

type ForbiddenError struct {
	BaseErrorResponse
}


func NewForbiddenError(detail string) CustomError {
	if detail == "" {
		detail = "Forbidden, you do not permissions to make that change."
	}

	return &InvalidCredentialsError{
		BaseErrorResponse{
			Type:     "https://shinypothos.com/docs/errors/forbidden",
			Title:    "Forbidden",
			Status:   http.StatusForbidden,
			Detail:   detail,
			Instance: "", // You can assign an instance value or leave it empty, based on your requirements.
		},
	}
}