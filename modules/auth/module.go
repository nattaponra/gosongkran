package auth

import (
	"github.com/nattaponra/gosongkran/core"
	"gorm.io/gorm"
)

type AuthModule struct {
	JwtSecret string
}

func (m *AuthModule) Name() string {
	return "auth"
}

func (m *AuthModule) Register(app *core.AppContext) {
	group := app.Router.Group("/auth")
	group.POST("/login", LoginHandler)
	group.POST("/register", RegisterHandler)
}

func (m *AuthModule) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
