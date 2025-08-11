package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rdcassin/quest-backend/internal/models"
	"go.uber.org/zap"
)

// Need to add the other relatiions.
type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	ImageURL string `json:"image_url"`
	Bio      string `json:"bio"`
	ClerkID  string `json:"clerk_id"`
}

func (a *App) CreateUserHandler(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		a.Logger.Warn("Invalid input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	user := models.User{
		ID:        uuid.NewString(),
		Username:  input.Username,
		Email:     input.Email,
		ImageURL:  input.ImageURL,
		Bio:       &input.Bio,
		ClerkID:   input.ClerkID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := a.DB.Create(&user).Error; err != nil {
		a.Logger.Error("Failed to create user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (a *App) GetUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := a.DB.First(&user, "id = ?", id).Error; err != nil {
		a.Logger.Warn("User not found", zap.String("id", id), zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
