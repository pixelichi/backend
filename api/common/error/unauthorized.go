package error

import "net/http"

type NotAuthorizedError struct {
	BaseErrorResponse
}


func NewNotAuthorizedError(detail string) CustomError {
	if detail == "" {
		detail = "Not Authorized"
	}

	return &BadRequestError{
		BaseErrorResponse{
			Type:     "https://pixelichi.com/docs/errors/not-authorized",
			Title:    "Not Authorized",
			Status:   http.StatusBadRequest,
			Detail:   detail,
			Instance: "", // You can assign an instance value or leave it empty, based on your requirements.
		},
	}
}