package util

import (
	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"` // "Err" diubah ke "Error" agar lebih konvensional
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
