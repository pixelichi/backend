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

type setProfilePictureResponse struct {
}

func SetProfilePicture(c *gin.Context) {

	tokenPayload := token.GetPayloadOrAbort(c)
	image := request_util.GetImageFromFormOrAbort(c, "file", image_util.ProfilePicConfig)
	rc := request_context.GetReqCtxOrInternalServerError(c)

	ostore_txn.UploadFileToUserDataOrAbort(c, rc.OS, tokenPayload.UserID, image, profPicFileName)

	c.Data(http.StatusOK, "image/jpeg", image.Bytes())
}
