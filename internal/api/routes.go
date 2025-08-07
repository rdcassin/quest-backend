package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rdcassin/quest-backend/internal/app"
)

func RegisterRoutes(router *gin.Engine, a *app.App) {
	router.GET("/healthz", a.HealthzHandler)
	api := router.Group("/api/v1")
	{
		api.GET("/ping", a.PingHandler)
		api.POST("/users", a.CreateUserHandler)
		api.GET("/users/:id", a.GetUserByIDHandler)
	}
}
