package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"shinypothos.com/api/common/error"
	db "shinypothos.com/db/sqlc"
	"shinypothos.com/util"
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


func SignUp(server *Server, ctx *gin.Context) {
	var req SignUpRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, error.NewBadRequestError(err.Error()))
		return
	}

	hashedPass, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.NewInternalServerError(err.Error()))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPass,
		Email:          req.Email,
	}

	user, err := server.Store.CreateUser(ctx, arg)

	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {
			code_name := pqErr.Code.Name()
			msg := fmt.Sprintf("Error Code: %v, Error Text: %v", code_name, err.Error())

			switch code_name {

			case "unique_violation":
				ctx.JSON(http.StatusForbidden, error.NewForbiddenError("User already exists, please use a different username and email."))
				return

			default:
				ctx.JSON(http.StatusInternalServerError, error.NewInternalServerError(msg))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, error.NewInternalServerError(err.Error()))
		return
	}

	// User successfully created, let's create a token
	accessToken, err := server.TokenMaker.CreateToken(user.ID, server.Config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, error.NewInternalServerError(err.Error()))
		return
	}

	// Set the token in an HttpOnly cookie
	httpOnlyCookie := http.Cookie{
		Name: "access_token",
		Value: accessToken,
		HttpOnly: true,
		Expires: time.Now().Add(12 * time.Hour),
		Path: "/",
	}

	http.SetCookie(ctx.Writer, &httpOnlyCookie)

	response := SignUpResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	ctx.JSON(http.StatusOK, response)
}
