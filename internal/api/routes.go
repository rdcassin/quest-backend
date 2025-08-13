package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rdcassin/quest-backend/internal/app"
	"github.com/rdcassin/quest-backend/internal/pkg/auth"
)

func RegisterRoutes(router *gin.Engine, a *app.App) {
	router.GET("/healthz", a.HealthzHandler)
	
	api := router.Group("/api/v1")
	{
		// Public endpoints
		api.GET("/ping", a.PingHandler)
		api.GET("/users", a.ListUsersHandler)
		api.GET("/users/:id", a.GetUserByIDHandler)

		// Protected endpoints
		protected := api.Group("/users")
		protected.Use(auth.ClerkAuthMiddleware())
		{
			protected.GET("/me", a.GetCurrentUserHandler)
			protected.POST("", a.CreateUserHandler)
			protected.PATCH("/:id", a.UpdateUserHandler)
			protected.DELETE("/:id", a.DeleteUserHandler)
		}
	}
}