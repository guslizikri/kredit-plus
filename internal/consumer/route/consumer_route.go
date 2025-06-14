package route

import (
	"sigmatech-kredit-plus/internal/consumer/handler"
	"sigmatech-kredit-plus/internal/consumer/repository"
	"sigmatech-kredit-plus/internal/consumer/usecase"
	"sigmatech-kredit-plus/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterConsumerRoutes(r *gin.Engine, db *sqlx.DB) {
	// Local DI
	repo := repository.NewConsumerRepository(db)
	usecase := usecase.NewConsumerUsecase(repo)
	handler := handler.NewConsumerHandler(usecase)
	consumer := r.Group("/consumers")
	{
		consumer.POST("/", middleware.MultiUploadMiddleware("consumer", []string{"photo_ktp", "photo_selfie"}), handler.CreateConsumer)
		consumer.GET("/:id", handler.GetConsumerByID)
	}
}
