package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"pixelichi.com/api/common/error"
	"pixelichi.com/token"
)

const (
	authorizationHeaderKey  = "authorization"
	validAuthorizationType  = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		// If auth header was not provided
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError(err.Error()))
			return
		}

		// Bearer {token}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError(err.Error()))

			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != validAuthorizationType {
			err := errors.New("unsupported authorization")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError(err.Error()))

			return
		}

		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			err := errors.New("invalid access token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError(err.Error()))

			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
