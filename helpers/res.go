package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type resHelper struct {
}

func NewResHelper() *resHelper {
	return &resHelper{}
}

func (r resHelper) ResponseSuccess(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": message,
		"data":    data,
	})
	return
}

func (r resHelper) ResponseFailed(c *gin.Context, err error, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"message": message,
		"errors":  err,
	})
}

func (r resHelper) ResponseValidationError(c *gin.Context, err error, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": message,
		"errors":  err,
	})
}

func (r resHelper) GenerateUuid() string {
	return uuid.New().String()
}
