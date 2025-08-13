package main

import (
	"log"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rdcassin/quest-backend/internal/api"
	"github.com/rdcassin/quest-backend/internal/app"
	"github.com/rdcassin/quest-backend/internal/database"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to init zap logger: %v", err)
	}
	defer func() { _ = logger.Sync() }()

	db := database.Connect()

	clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))

	application := &app.App{
		DB:     db,
		Logger: logger,
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	api.RegisterRoutes(router, application)

	logger.Info("Quest backend listening", zap.String("port", port))
	if err := router.Run(":" + port); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
