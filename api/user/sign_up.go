package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common/request_context"
	"shinypothos.com/api/common/request_util"
	"shinypothos.com/api/data/db_txn"
	db "shinypothos.com/db/sqlc"
	"shinypothos.com/util/password_util"
)

type SignUpRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type SignUpResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func SignUp(c *gin.Context) {
	rc, err := request_context.GetReqCtxOrInternalServerError(c)
	if err != nil {
		return
	}

	req, err := request_util.BindJSONOrAbort[SignUpRequest](c, &SignUpRequest{})
	if err != nil {
		return
	}

	hashedPass := password_util.HashPasswordOrAbort(c, req.Password)

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPass,
		Email:          req.Email,
	}

	user,err := db_txn.CreateUserOrAbort(c, rc.DB, &arg)
	if err != nil { // Could not create a user
		return
	}

	accessToken, err := (*rc.TokenMaker).CreateTokenOrAbort(c, user.ID, rc.Config.AccessTokenDuration)
	if err != nil {
		return
	}

	// Set the token in an HttpOnly cookie
	httpOnlyCookie := http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
		Expires:  time.Now().Add(rc.Config.AccessTokenDuration),
		SameSite: rc.Config.GetSameSite(),
		Path:     "/",
	}

	http.SetCookie(c.Writer, &httpOnlyCookie)

	response := SignUpResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	c.JSON(http.StatusOK, response)
}
