package ostore_txn

import (
	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common/request_context"
	"shinypothos.com/api/data/ostore"
	"shinypothos.com/token"
	"shinypothos.com/util/image_util"
)

const ProfPicFileName string = "profile_pic.jpg"

func UploadProfilePicOrAbort(c *gin.Context, os *ostore.OStore, userId int64, image *image_util.Image) {
	uploadFileToUserDataOrAbort(c, os, userId, image, ProfPicFileName)
}

func GetProfilePicPreSignedUrlOrAbort(c *gin.Context, os *ostore.OStore, userId int64) (string, error) {
	rc, err := request_context.GetReqCtxOrInternalServerError(c)
	if err != nil {
		return "", err
	}

	tokenPayload, err := token.GetPayloadOrAbort(c)
	if err != nil {
		return "", err
	}

	url, err := getPreSignedUrlForUserDataFileOrAbort(c, os, tokenPayload.UserID, ProfPicFileName, rc.Config.AccessTokenDuration)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}
