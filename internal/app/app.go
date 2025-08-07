package app

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Logger *zap.Logger
}
