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
	tokenPayload, err := token.GetPayloadOrAbort(c)
	if err != nil {
		return
	}

	rc,err := request_context.GetReqCtxOrInternalServerError(c)
	if err != nil {
		return
	}
	
	url,err := ostore_txn.GetProfilePicPreSignedUrlOrAbort(c,rc.OS,tokenPayload.UserID)
	if err != nil {
		return
	}

  response := GetProfilePictureResponse{
    ProfilePic: url,
	}

	c.JSON(http.StatusOK, response)
}
