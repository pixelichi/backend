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
	req, err := request_util.BindJSONOrAbort[loginUserRequest](c, &loginUserRequest{})
	if err != nil { //Unable to Bind Json to Variable - Bad Request
		return
	}

	rc, err := request_context.GetReqCtxOrInternalServerError(c)
	if err != nil { // User Context wasn't available
		return
	}

	user, err := db_fetch.GetUserOrAbort(c, rc.DB, req.Username)
	if err != nil { // User not found in DB
		return
	}

	password_util.CheckPasswordOrAbort(c, req.Password, user.HashedPassword)
	if err != nil { // Password was incorrect
		return
	}

	accessToken, err := (*rc.TokenMaker).CreateTokenOrAbort(c, user.ID, rc.Config.AccessTokenDuration)
	if err != nil { // Token could not be created for user
		return
	}

	// Set the token in an HttpOnly cookie
	httpOnlyCookie := http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
		Expires:  time.Now().Add(rc.Config.AccessTokenDuration),
		Path:     "/",
		SameSite: rc.Config.GetSameSite(),
	}

	http.SetCookie(c.Writer, &httpOnlyCookie)

	response := loginUserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	// Successful login, send back the response
	c.JSON(http.StatusOK, response)
}
