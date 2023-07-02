package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common/server_error"
	"shinypothos.com/token"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		accessTokenCookie, err := ctx.Request.Cookie("access_token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, server_error.NewNotAuthorizedError("No access token provided"))
			return
		}

		accessToken := accessTokenCookie.Value
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, server_error.NewNotAuthorizedError("Invalid Access Token"))
			return
		}

		ctx.Set(token.AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}