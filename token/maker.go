package token

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderKey  = "authorization"
	ValidAuthorizationType  = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

// Maker is an interface for managing tokens
type Maker interface {

	// Creates a new token for a specific username and duration
	CreateToken(userID int64, duration time.Duration) (string, error)

	CreateTokenOrAbort(ctx *gin.Context, userID int64, duration time.Duration) string

	// Check if the token is valid or not and if it is, it will return the payload in the token
	VerifyToken(token string) (Payload, error)
}
