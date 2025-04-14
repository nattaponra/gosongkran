package core

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type AppContext struct {
	Router *gin.Engine
	Config *AppConfig
	Logger *log.Logger
	DB     *gorm.DB
	Cache  *redis.Client
}
