package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pixelichi.com/api/common"
	"pixelichi.com/api/common/error"
	"pixelichi.com/token"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		accessTokenCookie, err := ctx.Request.Cookie("access_token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError("No access token provided"))
			return
		}

		accessToken := accessTokenCookie.Value
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError("Invalid Access Token"))
			return
		}

		ctx.Set(common.AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}

// func authMiddlewareWithAuthHeader(tokenMaker token.Maker) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

// 		// If auth header was not provided
// 		if len(authorizationHeader) == 0 {
// 			err := errors.New("authorization header is not provided")
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError(err.Error()))
// 			return
// 		}

// 		// Bearer {token}
// 		fields := strings.Fields(authorizationHeader)
// 		if len(fields) < 2 {
// 			err := errors.New("invalid authorization header format")
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError(err.Error()))

// 			return
// 		}

// 		authorizationType := strings.ToLower(fields[0])
// 		if authorizationType != validAuthorizationType {
// 			err := errors.New("unsupported authorization")
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError(err.Error()))

// 			return
// 		}

// 		accessToken := fields[1]

// 		payload, err := tokenMaker.VerifyToken(accessToken)
// 		if err != nil {
// 			err := errors.New("invalid access token")
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError(err.Error()))

// 			return
// 		}

// 		ctx.Set(authorizationPayloadKey, payload)
// 		ctx.Next()
// 	}
// }
