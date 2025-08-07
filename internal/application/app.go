package application

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func (a *App) HealthzHandler(c *gin.Context) {
	a.Logger.Info("Health check", zap.Time("time", time.Now()))
	c.JSON(200, gin.H{
		"message":   "ok",
		"timestamp": time.Now().Format(time.RFC3339),
		"database":  "connected",
	})
}

func (a *App) PingHandler(c *gin.Context) {
	a.Logger.Info("Ping endpoint hit")
	c.JSON(200, gin.H{"message": "pong"})
}

func (a *App) RegisterRoutes(router *gin.Engine) {
	router.GET("/healthz", a.HealthzHandler)
	api := router.Group("/api/v1")
	{
		api.GET("/ping", a.PingHandler)
	}
}
