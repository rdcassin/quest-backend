package app

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (a *App) HealthzHandler(c *gin.Context) {
	a.Logger.Info("Health check")
	c.JSON(200, gin.H{
		"message":   "ok",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

func (a *App) PingHandler(c *gin.Context) {
	a.Logger.Info("Ping endpoint hit")
	c.JSON(200, gin.H{"message": "pong"})
}
