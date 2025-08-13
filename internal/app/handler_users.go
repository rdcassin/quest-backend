package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rdcassin/quest-backend/internal/models"
	"go.uber.org/zap"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	ImageURL string `json:"image_url"`
	Bio      string `json:"bio"`
	ClerkID  string `json:"clerk_id"`
}

type UpdateUserInput struct {
	Username string  `json:"username,omitempty"`
	Email    string  `json:"email,omitempty"`
	ImageURL string  `json:"image_url,omitempty"`
	Bio      *string `json:"bio"`
	ClerkID  string  `json:"clerk_id,omitempty"`
}

func (a *App) CreateUserHandler(c *gin.Context) {
	input := CreateUserInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		a.Logger.Warn("Invalid input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	newUser := models.User{
		ID:        uuid.NewString(),
		Username:  input.Username,
		Email:     input.Email,
		ImageURL:  input.ImageURL,
		Bio:       &input.Bio,
		ClerkID:   input.ClerkID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := a.DB.Create(&newUser).Error; err != nil {
		a.Logger.Error("Failed to create user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (a *App) GetUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}
	if err := a.DB.First(&user, "id = ?", id).Error; err != nil {
		a.Logger.Warn("User not found", zap.String("id", id), zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (a *App) UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	input := UpdateUserInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		a.Logger.Warn("Invalid input", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	user := models.User{}
	if err := a.DB.First(&user, "id = ?", id).Error; err != nil {
		a.Logger.Warn("User not found", zap.String("id", id), zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if input.Username != "" {
		user.Username = input.Username
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.ImageURL != "" {
		user.ImageURL = input.ImageURL
	}
	if input.Bio != nil {
		user.Bio = input.Bio
	}
	if input.ClerkID != "" {
		user.ClerkID = input.ClerkID
	}
	user.UpdatedAt = time.Now()

	if err := a.DB.Save(&user).Error; err != nil {
		a.Logger.Error("Failed to update user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (a *App) ListUsersHandler(c *gin.Context) {
	users := []models.User{}
	if err := a.DB.Find(&users).Error; err != nil {
		a.Logger.Error("Failed to list users", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (a *App) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	if err := a.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		a.Logger.Error("Failed to delete user", zap.String("id", id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (a *App) GetCurrentUserHandler(c *gin.Context) {
	currentUserID, exists := c.Get("currentUserId")
	if !exists {
		a.Logger.Error("No user ID in context")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication context missing"})
		return
	}

	clerkID := currentUserID.(string)

	user := models.User{}
	if err := a.DB.First(&user, "clerk_id = ?", clerkID).Error; err != nil {
		a.Logger.Warn("Current user not found in database", zap.String("clerk_id", clerkID), zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{
			"error":    "User profile not found",
			"clerk_id": clerkID,
			"message":  "Please complete your profile setup",
		})
		return
	}

	a.Logger.Info("Current user retrieved", zap.String("user_id", user.ID), zap.String("username", user.Username))
	c.JSON(http.StatusOK, user)
}