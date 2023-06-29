package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"shinypothos.com/token"
)

const funcName = "set_profile_picture.go"

type request struct {

}

type response struct {
	Message string `json:"message"`
}

func SetProfilePicture(ctx *gin.Context) {
	// var req request

	// Bind the request body to the request struct””
	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	// Return an error if the request body is invalid
	// 	ctx.JSON(http.StatusBadRequest, error.NewBadRequestError("Invalid Request Body - "+err.Error()))
	// 	return
	// }

	userCtx, err := token.GetUserPayloadOrFatal(ctx)
	if err != nil {
		log.Fatalf("%v - Unable to get user information from auth token, %v", funcName, err)
		return
  }

  log.Printf("This is the user we are dealing with user id: %v", userCtx.UserID)

	response := response{
		Message: strconv.FormatInt(userCtx.UserID, 10),
	}

	// Successful login, send back the response
	ctx.JSON(http.StatusOK, response)
}
