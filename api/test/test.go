package test

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"shinypothos.com/api/minio"
)

type testResponse struct {
	Msg string `json:"msg"`
}


func Minio(ctx *gin.Context) {

	bucket, err := minio.MinioClient.ListBuckets(ctx)
  if err != nil{
    fmt.Printf("err: %v\n", err)  
  }

  fmt.Printf("bucket: %v\n", bucket)

	response := testResponse{
		Msg: "test",
	}

	ctx.JSON(http.StatusOK, response)
}