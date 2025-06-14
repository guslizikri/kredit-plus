package middleware

import (
	"net/http"
	"strings"

	"sigmatech-kredit-plus/pkg"

	"github.com/gin-gonic/gin"
)

func Auth(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid header type"})
			ctx.Abort()
			return
		}

		tokenStr := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
		check, err := pkg.VerifyToken(tokenStr)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			ctx.Abort()
			return
		}

		var valid bool
		for _, r := range role {
			if r == check.Role {
				valid = true
				break
			}
		}

		if !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "you not have permission"})
			ctx.Abort()
			return
		}

		ctx.Set("consumerId", check.ConsumerId)
		ctx.Set("adminId", check.AdminId)
		ctx.Next()
	}
}
