package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common"
	"shinypothos.com/api/common/request_context"
	"shinypothos.com/api/common/server_error"
)

type Server = common.Server

const accessTokenKey = "access_token"


type CheckAuthResponse struct {
	Message string `json:"message"`
}


func accessTokenOrInternalServerError(c *gin.Context) (string,error) {
	accessTokenCookie, err := c.Request.Cookie(accessTokenKey)
	if err != nil { // Could not get access token from cookie
		c.AbortWithStatusJSON(http.StatusUnauthorized, server_error.NewNotAuthorizedError("No Access Token Provided"))
		return "", err
	}

	accessToken := accessTokenCookie.Value
	return accessToken, nil
}


func CheckAuth(c *gin.Context) {
	reqCtx, err := request_context.GetReqCtxOrInternalServerError(c)
	if err != nil { // Could not get request context
		return
	}

	accessToken, err := accessTokenOrInternalServerError(c)
	if err != nil { // Could not get access token from cookie 
		return
	}

	_, err = (*reqCtx.TokenMaker).VerifyToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, server_error.NewNotAuthorizedError("Invalid Access Token"))
		return
	}

	response := CheckAuthResponse{
		Message: "Access granted",
	}

	c.JSON(http.StatusOK, response)
}
