package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "os"
    "time"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    router := gin.New()

    // Logging middleware
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    // Health check endpoint
    router.GET("/healthz", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "ok", "timestamp": time.Now().Format(time.RFC3339)})
    })

    // Versioned API (example ping route)
    api := router.Group("/api/v1")
    {
        api.GET("/ping", func(c *gin.Context) {
            c.JSON(200, gin.H{"message": "pong"})
        })
    }

    log.Printf("Quest backend listening on :%s", port)
    if err := router.Run(":" + port); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}