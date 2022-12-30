package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "pixelichi.com/db/sqlc"
	"pixelichi.com/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPass, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPass,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {

			code_name := pqErr.Code.Name()
			msg := fmt.Sprintf("Error Code: %v, Error Text: %v", code_name, err.Error())

			switch code_name {

			case "unique_violation":
				ctx.JSON(http.StatusForbidden, msg)
				return

			default:
				ctx.JSON(http.StatusInternalServerError, msg)
				return
			}

		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// don't respond with password or hashed pass
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponseStruct struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	// Unfurtling the request to login
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Fetch user that is attempting to log in from the DB
	user, err := server.store.GetUserFromUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("invalid username")))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Check if password provided by the client is correct or not
	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("invalid credentials")))
		return
	}

	// Password was correct, lets create a token
	token, err := server.tokenMaker.CreateToken(user.ID, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Successful login, send back the response
	ctx.JSON(http.StatusOK, loginUserResponseStruct{
		AccessToken: token,
		User:        newUserResponse(user),
	})
}
