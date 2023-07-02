package ostore_txn

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common/server_error"
	"shinypothos.com/api/data/ostore"
	"shinypothos.com/util/image_util"
)

const UserDataBucket string = "shinypothos-user-data"

// func GetProfilePicUrl(c context.Context, s *common.Server) (string , error){
//   (*s.ObjectStore).GetSignedUrlForUserDataFile(c,)
// }

func getUserDataObjectPrefix(userID int64) string {
	return strconv.FormatInt(userID, 10) + "/"
}

func UploadFileToUserDataOrAbort(c *gin.Context, os *ostore.OStore, userId int64, image image_util.Image, objectName string) {
	err := (*os).UploadFile(c, UserDataBucket, getUserDataObjectPrefix(userId)+objectName, image.Reader, int64(image.Size))
	
  if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError("Unable to upload file."))
	}
}


func GetSignedUrlForUserDataFileOrAbort(c *gin.Context, os *ostore.OStore, object string, expDuration time.Duration, ) (url.URL) {

	presignedURL, err := (*os).GetSignedUrlForUserDataFileEmptyParam(c, UserDataBucket, "myobject", expDuration)
	if err != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError("Could not get pre-signed url for file."))
	}

	return presignedURL
}