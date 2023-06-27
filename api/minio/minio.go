package minio

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinioClient(endpoint, accessKey, secretKey string) {
	var err error
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln("Failed to initialize MinIO client:", err)
	}
}

func CreateBucketIfDoesntExist(ctx context.Context, bucketName string) {
  // Create a new bucket if it doesn't exist
	exists, err := MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
    log.Fatalf("CreateBucketIfDoesntExist: Error Checking if bucket exists - %v\n", err)
	}

	if !exists {
		err := MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("CreateBucketIfDoesntExist: Error creating bucket - %v\n", err)
		}
	}
}