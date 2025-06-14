package route

import (
	"sigmatech-kredit-plus/internal/auth/handler"
	"sigmatech-kredit-plus/internal/auth/repository"
	"sigmatech-kredit-plus/internal/auth/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterAuthRoutes(r *gin.Engine, db *sqlx.DB) {
	// Local DI
	repo := repository.NewAuthRepository(db)
	usecase := usecase.NewAuthUsecase(repo)
	handler := handler.NewAuthHandler(usecase)
	auth := r.Group("/auth")
	{
		auth.POST("/consumer-login", handler.ConsumerLogin)
		// for now logic enpoint admin login still hardcode
		auth.POST("/admin-login", handler.AdminLogin)
	}
}
