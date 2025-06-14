package route

import (
	"sigmatech-kredit-plus/internal/middleware"
	"sigmatech-kredit-plus/internal/user/handler"
	"sigmatech-kredit-plus/internal/user/repository"
	"sigmatech-kredit-plus/internal/user/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterUserRoutes(r *gin.Engine, db *sqlx.DB) {
	// Local DI
	repo := repository.NewUserRepository(db)
	usecase := usecase.NewUserUsecase(repo)
	handler := handler.NewUserHandler(usecase)
	user := r.Group("/users")
	{
		user.POST("/", middleware.MultiUploadMiddleware("user", []string{"photo_ktp", "photo_selfie"}), handler.CreateUser)
		user.GET("/:id", handler.GetUserByID)
	}
}
