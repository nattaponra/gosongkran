package core

import (
	"log"

	"github.com/gin-gonic/gin"
)

type AppContext struct {
	Router *gin.Engine
	Config AppConfig
	Logger *log.Logger
	// DB     *gorm.DB     // If using DB
	// Cache  *redis.Client // Optional
}
type AppConfig struct {
	JWTSecret string
	AppEnv    string
	Port      string
	// Add more config values here
}
