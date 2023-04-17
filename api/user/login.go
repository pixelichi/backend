package user

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common"
	"shinypothos.com/api/common/error"
	"shinypothos.com/util"
)

type Server = common.Server

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}


func LoginUser(server *Server, ctx *gin.Context) {
	var req loginUserRequest

	// Bind the request body to the request struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// Return an error if the request body is invalid
		ctx.JSON(http.StatusBadRequest, error.NewBadRequestError("Invalid Request Body - " + err.Error()))
		return
	}

	// Fetch user that is attempting to log in from the DB
	user, err := server.Store.GetUserFromUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, error.NewInvalidCredentialsError("Couldn't find user - " + req.Username))
			return
		}

		ctx.JSON(http.StatusInternalServerError, error.NewInternalServerError(err.Error()))
		return
	}

	// Check if password provided by the client is correct or not
	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, error.NewInvalidCredentialsError("Password provided is incorrect. Try again."))
		return
	}

	// Password was correct, lets create a token
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

	response := loginUserResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}

	// Successful login, send back the response
	ctx.JSON(http.StatusOK, response)
}