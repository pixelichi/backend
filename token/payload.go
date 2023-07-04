package token

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"shinypothos.com/api/common/server_error"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    int64     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// creates a new token payload with a specific username and duration
func NewPayload(userId int64, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		UserID:    userId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func GetPayloadOrAbort(c *gin.Context) (*Payload, error) {

	auth_info, exists := c.Get(AuthorizationPayloadKey)
	if !exists {
		err := server_error.NewNotAuthorizedError("User Authentication information was not available in the request context.")
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return nil, err
	}

	payload, ok := auth_info.(Payload)
	if !ok {
		err := server_error.NewInternalServerError("Could not read user information")
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return nil, err
	}

	return &payload, nil
}
