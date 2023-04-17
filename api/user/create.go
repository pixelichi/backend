package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"shinypothos.com/api/common/error"
	db "shinypothos.com/db/sqlc"
	"shinypothos.com/util"
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


func CreateUser(server *Server, ctx *gin.Context) {
	var req createUserRequest

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
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.Store.CreateUser(ctx, arg)
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

		ctx.JSON(http.StatusInternalServerError, error.NewInternalServerError(err.Error()))
		return
	}

	response := userResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
	}

	ctx.JSON(http.StatusOK, response)
}
