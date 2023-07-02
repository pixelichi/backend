package image_util

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
	"shinypothos.com/api/common/server_error"
	"shinypothos.com/util/conv"
)

type Image struct {
	Reader io.Reader
	Size uint64
}

func (image *Image) Bytes()[]byte{
	  buf := new(bytes.Buffer)
    buf.ReadFrom(image.Reader)
    return buf.Bytes()
}
 
var emptyImage = Image{}

type ImageConfig struct {
    Width  uint
    Height uint
    Quality  uint8 // Must be <= 0 && < 100
		MaxSizeMB uint8
}

var ProfilePicConfig = ImageConfig{
	Width: 360,
	Height: 360,
	Quality: 100,
	MaxSizeMB: 5,
}

func ImageToJPEGWithConfigOrAbort(c *gin.Context, fileHeader *multipart.FileHeader, config ImageConfig) Image {
	image, err := imageToJPEGWithConfig(fileHeader, config)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, server_error.NewInternalServerError("Could not convert image to jpeg. Please only upload .jpg .jpeg or .png"))
	}

	return image
}

func imageToJPEGWithConfig(file *multipart.FileHeader, config ImageConfig)  (Image, error){
	return ImageToJPEG(file, config.MaxSizeMB, config.Width, config.Height, uint8(config.Quality))
}

func ImageToJPEG(file *multipart.FileHeader, maxFileSizeMB uint8, width uint, height uint, quality uint8) (Image, error) {
  var maxFileSizeBytes int64 = 0

	if maxFileSizeMB == 0 { // Use default 5 MB
		maxFileSizeMB = 5
  }

	// Convert from MB to Bytes
	maxFileSizeBytes = conv.MbUint8ToBytes(maxFileSizeMB)

	// Check the file size
	if file.Size > maxFileSizeBytes {
		return emptyImage, errors.New("File size is > " + strconv.FormatUint(uint64(maxFileSizeMB), 10) + " MB")
	}

	// Check the file content type
	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image/") {
		return emptyImage, errors.New("File is not an image according to it's header")
	}

	// Read file contents
	src, err := file.Open()
	if err != nil {
		return emptyImage, errors.New("Unable to open() file")
	}
	defer src.Close()

	// The converted image will be written to this buffer
	var target bytes.Buffer
	err = ConvertToJPEG(src, &target, width, height, quality)
	if err != nil {
		return emptyImage, errors.New("Something went wrong when converting the image to JPEG")
	}

	// Return image!
	return Image{
		Reader: &target,
		Size: uint64(target.Len()),
		}, nil
}

// If you want to maintain aspect ratio you can simply pas a 0 to one of the [width, height] parameters
func ConvertToJPEG(src io.Reader, target io.Writer, width uint, height uint, quality uint8) error {
	if quality < 1 {
		quality = 1
	}
	if quality > 100 {
		quality = 100
	}

	img, _, err := image.Decode(src)
	if err != nil {
		return err
	}

	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	options := jpeg.Options{
		Quality: int(quality) ,
	}

	return jpeg.Encode(target, resizedImg, &options)
}

