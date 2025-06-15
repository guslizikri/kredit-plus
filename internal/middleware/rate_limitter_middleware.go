package middleware

import (
	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
)

// NewRateLimiterMiddleware creates a rate limiter middleware with the given format, e.g., \"5-M\" (5 per minute)
func RateLimiterMiddleware(rate limiter.Rate) gin.HandlerFunc {
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	middleware := ginlimiter.NewMiddleware(instance)

	return middleware
}
