package core

import (
	"log"

	"github.com/gin-gonic/gin"
)

type AppContext struct {
	Router *gin.Engine
	Config map[string]interface{} // Can use viper or your own config
	Logger *log.Logger
	// DB     *gorm.DB     // If using DB
	// Cache  *redis.Client // Optional
}
