package route

import (
	"sigmatech-kredit-plus/internal/limit/handler"
	"sigmatech-kredit-plus/internal/limit/repository"
	"sigmatech-kredit-plus/internal/limit/usecase"
	"sigmatech-kredit-plus/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterLimitRoutes(r *gin.Engine, db *sqlx.DB) {
	// Local DI
	repo := repository.NewLimitRepository(db)
	usecase := usecase.NewLimitUsecase(repo)
	handler := handler.NewLimitHandler(usecase)
	limit := r.Group("/limits")
	{
		limit.POST("/set-limit/:consumerId", middleware.Auth("admin"), handler.SetLimit)
	}
}
