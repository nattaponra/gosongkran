package auth

import "github.com/nattaponra/gosongkran/core"

type AuthModule struct{}

func (m *AuthModule) Name() string {
	return "auth"
}

func (m *AuthModule) Register(app *core.AppContext) {
	group := app.Router.Group("/auth")
	group.POST("/login", LoginHandler)
	group.POST("/register", RegisterHandler)
}

func init() {
	core.RegisterModule(&AuthModule{}) // Auto-register during package init
}
