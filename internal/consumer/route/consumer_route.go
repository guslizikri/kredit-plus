package route

import (
	"sigmatech-kredit-plus/internal/consumer/handler"
	"sigmatech-kredit-plus/internal/consumer/repository"
	"sigmatech-kredit-plus/internal/consumer/usecase"
	limit_repo "sigmatech-kredit-plus/internal/limit/repository"
	"sigmatech-kredit-plus/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterConsumerRoutes(r *gin.Engine, db *sqlx.DB) {
	// Local DI
	repo := repository.NewConsumerRepository(db)
	limitRepo := limit_repo.NewLimitRepository(db)
	usecase := usecase.NewConsumerUsecase(repo, limitRepo)
	handler := handler.NewConsumerHandler(usecase)
	consumer := r.Group("/consumers")
	{
		consumer.POST("/", middleware.MultiUploadMiddleware("consumer", []string{"photo_ktp", "photo_selfie"}), handler.CreateConsumer)
		consumer.GET("/:id", middleware.Auth("consumer", "admin"), handler.GetConsumerByID)
	}
}
