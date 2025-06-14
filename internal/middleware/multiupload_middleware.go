package middleware

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func MultiUploadMiddleware(folderName string, attributes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		basePath := filepath.Join("../../kp-files", folderName)
		if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory: " + err.Error()})
			return
		}

		form, err := c.MultipartForm()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid multipart form: " + err.Error()})
			return
		}

		for _, attribute := range attributes {
			files := form.File[attribute]

			// fallback jika multipart tidak punya key
			if len(files) == 0 {
				file, err := c.FormFile(attribute)
				if err != nil {
					if errors.Is(err, http.ErrMissingFile) {
						continue
					}
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error retrieving the file: " + err.Error()})
					return
				}
				files = []*multipart.FileHeader{file}
			}

			var savedPaths []string
			for _, file := range files {
				if !isImage(file.Filename) {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": fmt.Sprintf("Invalid file type for field %s: %s", attribute, file.Filename),
					})
					return
				}

				filename := fmt.Sprintf("%d-%s-%s", time.Now().UnixNano(), attribute, sanitizeFilename(file.Filename))
				fullPath := filepath.Join(basePath, filename)

				if err := c.SaveUploadedFile(file, fullPath); err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
						"error": "Failed to save file: " + err.Error(),
					})
					return
				}

				relativePath := fmt.Sprintf("/%s/%s", folderName, filename)
				savedPaths = append(savedPaths, relativePath)
			}

			// Simpan string jika 1 file, []string jika lebih
			if len(savedPaths) == 1 {
				c.Set(attribute, savedPaths[0])
			} else {
				c.Set(attribute, savedPaths)
			}
		}

		c.Next()
	}
}

func isImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp":
		return true
	default:
		return false
	}
}

func sanitizeFilename(filename string) string {
	return strings.ReplaceAll(filename, " ", "_")
}
