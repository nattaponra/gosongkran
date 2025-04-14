package user

import (
	"github.com/gin-gonic/gin"
)

func GetProfileHandler(c *gin.Context) {
	user := c.MustGet("user") // token claims
	c.JSON(200, gin.H{"profile": user})
}
