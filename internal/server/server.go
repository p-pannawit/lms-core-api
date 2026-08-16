package server

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/p-pannawit/salung-api/internal/health"
)

func New(database *sql.DB) *gin.Engine {
	healthRepository := health.NewPostgresRepository(database)
	healthService := health.NewService(healthRepository)
	healthHandler := health.NewHandler(healthService)

	router := gin.Default()
	router.GET("/health", healthHandler.Check)

	apiV1 := router.Group("/api/v1")
	apiV1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	return router
}
