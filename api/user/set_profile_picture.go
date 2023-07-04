package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common/request_context"
	"shinypothos.com/api/common/request_util"
	"shinypothos.com/api/data/ostore_txn"
	"shinypothos.com/token"
	"shinypothos.com/util/image_util"
)

func SetProfilePicture(c *gin.Context) {
	tokenPayload, err := token.GetPayloadOrAbort(c)
	if err != nil {
		return
	}

	image,err := request_util.GetImageFromFormOrAbort(c, "file", image_util.ProfilePicConfig)
	if err != nil {
		return 
	}

	rc,err := request_context.GetReqCtxOrInternalServerError(c)
	if err != nil {
		return
	}

	ostore_txn.UploadProfilePicOrAbort(c, rc.OS, tokenPayload.UserID, image)

	c.Data(http.StatusOK, "image/jpeg", image.Bytes())
}
