package user

import (
	"github.com/gin-gonic/gin"
)

type Module struct{}

func (m *Module) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/users")
	group.GET("/me", GetProfileHandler)

}
