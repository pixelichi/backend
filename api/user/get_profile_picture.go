package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common/request_context"
	"shinypothos.com/api/data/ostore_txn"
	"shinypothos.com/token"
)

type GetProfilePictureResponse struct {
	ProfilePic string `json:"profile_pic"`
}

func GetProfilePicture(c *gin.Context) {
	tokenPayload := token.GetPayloadOrAbort(c)
	rc := request_context.GetReqCtxOrInternalServerError(c)
	
	url := ostore_txn.GetProfilePicPreSignedUrlOrAbort(c,rc.OS,tokenPayload.UserID)

  response := GetProfilePictureResponse{
    ProfilePic: url,
	}

	c.JSON(http.StatusOK, response)
}
