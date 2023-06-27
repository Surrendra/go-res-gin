package helpers

import (
	"github.com/gin-gonic/gin"
)

type fileSystemHelper struct {
}

func NewFileSystemHelper() *fileSystemHelper {
	return &fileSystemHelper{}
}

type FileSystemHelper interface {
	UploadFile(c *gin.Context, filename string, path string) (string, error)
}

func (h fileSystemHelper) UploadFile(c *gin.Context, filename string, path string) (string, error) {
	photo, errPhoto := c.FormFile("photo")
	if errPhoto != nil {
		return "", errPhoto
	}
	errSaveUploadFile := c.SaveUploadedFile(photo, path)
	if errSaveUploadFile != nil {
		return "", errSaveUploadFile
	}
	return path, nil
}
