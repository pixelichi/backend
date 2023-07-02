package user

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"shinypothos.com/api/common/error"
// 	"shinypothos.com/api/minio"
// 	"shinypothos.com/token"
// 	"shinypothos.com/util/image"
// )

// type getProfileResponse struct {
// 	ProfilePicUrl    string `json:"profile_pic_url"`
// }

// func GetProfile(ctx *gin.Context) {

// 	userCtx, err := token.GetUserPayloadOrFatal(ctx)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.NewNotAuthorizedError("Error Setting Profile Picture, unauthorized"))
// 		return
// 	}

// 	err = minio.UploadFileToUserData(ctx, userCtx.UserID, image, profPicFileName)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
// 	}

// 	ctx.Data(http.StatusOK, "image/jpeg", image.Bytes())
// 	response := response{}

// 	// Successful login, send back the response
// 	ctx.JSON(http.StatusOK, response)
// }
