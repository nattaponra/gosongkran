package user

import (
	"github.com/nattaponra/gosongkran/core"
	"github.com/nattaponra/gosongkran/modules/auth"
)

type UserModule struct{}

func (m *UserModule) Name() string {
	return "user"
}

func (m *UserModule) Register(app *core.AppContext) {
	group := app.Router.Group("/users")

	// Protected route
	group.GET("/me", auth.JWTMiddleware(&app.Config), GetProfileHandler)
}
