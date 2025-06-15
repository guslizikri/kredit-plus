package route

import (
	"sigmatech-kredit-plus/internal/auth/handler"
	"sigmatech-kredit-plus/internal/auth/repository"
	"sigmatech-kredit-plus/internal/auth/usecase"
	"sigmatech-kredit-plus/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/ulule/limiter/v3"
)

func RegisterAuthRoutes(r *gin.Engine, db *sqlx.DB) {
	// Local DI
	repo := repository.NewAuthRepository(db)
	usecase := usecase.NewAuthUsecase(repo)
	handler := handler.NewAuthHandler(usecase)

	rate, _ := limiter.NewRateFromFormatted("5-M")
	auth := r.Group("/auth")
	auth.Use(middleware.RateLimiterMiddleware(rate))
	{
		auth.POST("/consumer-login", handler.ConsumerLogin)
		// for now logic enpoint admin login still hardcode
		auth.POST("/admin-login", handler.AdminLogin)
	}
}
