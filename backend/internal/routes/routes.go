package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"aeb-project/backend/internal/handlers"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {

	healthHandler := handlers.NewHealthHandler(db)

	api := router.Group("/api")
	{
		api.GET("/health", healthHandler.Check)
	}

}
