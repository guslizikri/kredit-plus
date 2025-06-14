package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	consumerRoute "sigmatech-kredit-plus/internal/consumer/route"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Set up CORS options
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply the CORS middleware to the router
	router.Use(cors.New(corsConfig))

	consumerRoute.RegisterConsumerRoutes(router, db)

	return router
}
