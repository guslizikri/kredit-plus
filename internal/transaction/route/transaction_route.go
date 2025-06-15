package route

import (
	"sigmatech-kredit-plus/internal/common"
	limit_repository "sigmatech-kredit-plus/internal/limit/repository"
	"sigmatech-kredit-plus/internal/middleware"
	"sigmatech-kredit-plus/internal/transaction/handler"
	"sigmatech-kredit-plus/internal/transaction/repository"
	"sigmatech-kredit-plus/internal/transaction/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterTransactionRoutes(r *gin.Engine, db *sqlx.DB) {
	// Local DI
	repo := repository.NewTransactionRepository(db)
	limitRepo := limit_repository.NewLimitRepository(db)
	trxManager := common.NewTransactionManager(db)
	usecase := usecase.NewTransactionUsecase(trxManager, limitRepo, repo)
	handler := handler.NewTransactionHandler(usecase)
	transaction := r.Group("/transactions")
	{
		transaction.POST("/", middleware.Auth("consumer"), handler.CreateTransaction)
		transaction.GET("/consumer-histories", middleware.Auth("consumer", "admin"), handler.GetTransactionHistory)
	}
}
