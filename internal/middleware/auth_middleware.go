package middleware

import (
	"net/http"
	"sigmatech-kredit-plus/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(role ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var valid bool
		var header string

		header = ctx.GetHeader("Authorization")
		if header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			ctx.Abort()
			return
		}

		if !strings.Contains(header, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid header type"})
			ctx.Abort()
			return
		}

		tokens := strings.TrimPrefix(header, "Bearer ")
		tokens = strings.TrimSpace(tokens)

		check, err := pkg.VerifyToken(tokens)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		for _, r := range role {
			if r == check.Role {
				valid = true
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
