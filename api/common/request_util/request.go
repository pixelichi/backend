package request_util

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/common/server_error"
	"shinypothos.com/util/image_util"
)

// obj should be a &struct{}, returning value is a POINTER! *struct{} type.
// eg. BindJSONOrAbort(c, &loginUserRequest{}).(*loginUserRequest)
func BindJSONOrAbort(c *gin.Context, obj any) any {
	err := c.ShouldBindJSON(obj)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, server_error.NewBadRequestError("Invalid Request Body - "+err.Error()))
	}

	return obj
}


func GetFileHeaderFromFormOrAbort(c *gin.Context, formKey string) *multipart.FileHeader {
	header, err := c.FormFile(formKey)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, server_error.NewBadRequestError("Didn't receive "+formKey+" in request"))
	}

	return header
}


func GetImageFromFormOrAbort(c *gin.Context, formKey string, imageConfig image_util.ImageConfig) image_util.Image {
	fHeader := GetFileHeaderFromFormOrAbort(c, formKey)
	image := image_util.ImageToJPEGWithConfigOrAbort(c, fHeader, image_util.ProfilePicConfig)

	return image
}