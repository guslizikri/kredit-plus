package util

import (
	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// SendResponse is a helper to send JSON responses in Gin
func SendResponse(c *gin.Context, statusCode int, data interface{}, message string) {
	resp := HttpResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, resp)
}

// SendResponseWithMeta adds a Meta field (e.g., for pagination)
func SendResponseWithMeta(c *gin.Context, statusCode int, data, meta interface{}, message string) {
	resp := HttpResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	c.JSON(statusCode, resp)
}

// SendResponseWithError sends JSON with an error field
func SendResponseWithError(c *gin.Context, statusCode int, data interface{}, message string, err interface{}) {
	resp := HttpResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
		Error:   err,
	}
	c.JSON(statusCode, resp)
}
