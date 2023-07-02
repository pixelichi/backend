package ostore

import (
	"errors"
	"io"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"shinypothos.com/util/conv"
)

// const (
// 	maxFileSizeMB  uint8  = 10
// )

// Minio is a ObjectStore
type Minio struct {
	client *minio.Client
}

var maxFileSizeBytes int64 = conv.MbUint8ToBytes(maxFileSizeMB)

func NewMinioObjectStore(endpoint string, accessKey string, secretKey string) OStore {
	newMinioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatal("cannot create object_store:", err)
	}

	minio := &Minio{
		client: newMinioClient,
	}

	// minio.CreateBucketIfDoesntExist(context.Background(), user_data.UserDataBucket)

	return minio
}

func (m *Minio) createBucketIfDoesntExist(ctx *gin.Context, bucketName string) {
	// Create a new bucket if it doesn't exist
	exists, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("CreateBucketIfDoesntExist: Error Checking if bucket exists - %v\n", err)
	}

	if !exists {
		err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("CreateBucketIfDoesntExist: Error creating bucket - %v\n", err)
		}
	}
}

func (m *Minio) UploadFile(ctx *gin.Context, bucketName, objectName string, reader io.Reader, fileSize int64) error {

	// Validate file size to not exceed max-size in MB
	if fileSize > maxFileSizeBytes {
		return errors.New("File size is > " + strconv.FormatUint(uint64(maxFileSizeMB), 10) + " MB")
	}

	_, err := m.client.PutObject(ctx, bucketName, objectName, reader, fileSize, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (m *Minio) GetSignedUrlForUserDataFile(c *gin.Context, bucket string, object string, expDuration time.Duration, reqParams url.Values) (url.URL, error) {

	presignedURL, err := m.client.PresignedGetObject(c, bucket, object, expDuration, reqParams)
	if err != nil {
		return url.URL{}, err
	}

	return *presignedURL, nil
}

func (m *Minio) GetSignedUrlForUserDataFileEmptyParam(c *gin.Context, bucket string, object string, expDuration time.Duration) (url.URL, error) {
	reqParams := make(url.Values)

	presignedURL, err := m.client.PresignedGetObject(c, bucket, object, expDuration, reqParams)
	if err != nil {
		return url.URL{}, err
	}

	return *presignedURL, nil
}
