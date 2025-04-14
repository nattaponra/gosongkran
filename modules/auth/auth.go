package auth

import "github.com/gin-gonic/gin"

type Module struct{}

func (m *Module) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/auth")
	group.POST("/login", LoginHandler)
	group.POST("/register", RegisterHandler)
}
