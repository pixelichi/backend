package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common"
	"shinypothos.com/api/common/request_context"
	"shinypothos.com/api/common/request_util"
	db_fetch "shinypothos.com/api/data/db_txn"
	"shinypothos.com/util/password_util"
)

type Server = common.Server

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}


func LoginUser(c *gin.Context) {
	req := request_util.BindJSONOrAbort(c, &loginUserRequest{}).(*loginUserRequest)
	rc := request_context.GetReqCtxOrInternalServerError(c)

	user := db_fetch.GetUserOrAbort(c, rc.DB, req.Username)
	password_util.CheckPasswordOrAbort(c, req.Password, user.HashedPassword)
	accessToken := (*rc.TokenMaker).CreateTokenOrAbort(c, user.ID, rc.Config.AccessTokenDuration)


	// Set the token in an HttpOnly cookie
	httpOnlyCookie := http.Cookie{
		Name: "access_token",
		Value: accessToken,
		HttpOnly: true,
		Expires: time.Now().Add(rc.Config.AccessTokenDuration),
		Path: "/",
	}

	http.SetCookie(c.Writer, &httpOnlyCookie)

	response := loginUserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	// Successful login, send back the response
	c.JSON(http.StatusOK, response)
}