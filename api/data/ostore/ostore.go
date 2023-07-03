package ostore

import (
	"io"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	maxFileSizeMB  uint8  = 10
)

// OStore is an interface for integrating with s3 compliant object stores
type OStore interface {
	createBucketIfDoesntExist(c *gin.Context, bucketName string)
  UploadObject(c *gin.Context, bucketName, objectName string, reader io.Reader, fileSize int64) error
	GetSignedUrlForObject(c *gin.Context, bucket string, object string, expDuration time.Duration, reqParams url.Values) (url.URL, error)
}
