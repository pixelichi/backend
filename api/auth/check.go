package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pixelichi.com/api/common"
	"pixelichi.com/api/common/error"
)

type Server = common.Server

type CheckAuthResponse struct {
	Message string `json:"message"`
}


func CheckAuth(server *Server, ctx *gin.Context) {

	accessTokenCookie, err := ctx.Request.Cookie("access_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, error.NewNotAuthorizedError("No Access Token Provided"))
		return
	}

	accessToken := accessTokenCookie.Value

	// Password was correct, lets create a token
	_ , err = server.TokenMaker.VerifyToken(accessToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, error.NewNotAuthorizedError("Invalid Access Token"))
		return
	}
	
	response := CheckAuthResponse{
		Message: "Access granted",
	}

	ctx.JSON(http.StatusOK, response)
}