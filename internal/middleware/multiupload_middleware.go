package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func MultiUploadMiddleware(folderName string, attributes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, attribute := range attributes {
			fileHeader, err := c.FormFile(attribute)
			if err != nil {
				if errors.Is(err, http.ErrMissingFile) || fileHeader == nil {
					continue
				}
				c.AbortWithStatusJSON(400, gin.H{"error": "Error retrieving the file: " + err.Error()})
				return
			}

			absPath, _ := filepath.Abs(fmt.Sprintf("../../kp-files/%s", folderName))

			if err = os.MkdirAll(absPath, os.ModePerm); err != nil {
				c.AbortWithStatusJSON(500, gin.H{"error": "Error creating directory: " + err.Error()})
				return
			}

			filename := fmt.Sprintf("%d-%s.%s", time.Now().Unix(), attribute, fileHeader.Filename)
			outputFilePath := filepath.Join(absPath, filename)

			if err := c.SaveUploadedFile(fileHeader, outputFilePath); err != nil {
				c.AbortWithStatusJSON(500, gin.H{"error": "Error saving file: " + err.Error()})
				return
			}

			filePath := fmt.Sprintf("/%s/%s", folderName, filename)
			c.Set(attribute, filePath)
		}

		c.Next()
	}
}
