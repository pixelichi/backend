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

const (
	UserDataBucket string = "shinypothos-user-data"
)

func getUserDataObjectPrefix(userID int64) string {
	return strconv.FormatInt(userID, 10) + "/"
}

func uploadFileToUserDataOrAbort(c *gin.Context, os *ostore.OStore, userId int64, image *image_util.Image, objectName string) {
	err := (*os).UploadObject(c, UserDataBucket, getUserDataObjectPrefix(userId)+objectName, image.Reader, int64(image.Size))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError("Unable to upload file."))
	}
}

func getPreSignedUrlForUserDataFileOrAbort(c *gin.Context, os *ostore.OStore, userId int64, object string, expDuration time.Duration) (*url.URL, error) {
	presignedURL, err := (*os).GetSignedUrlForObject(c, UserDataBucket, strconv.FormatInt(userId, 10) + "/" + object, expDuration, url.Values{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError("Could not get pre-signed url for file."))
		return nil, err
	}

	return &presignedURL, nil
}
