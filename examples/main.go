package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nattaponra/gosongkran/core"
	"github.com/nattaponra/gosongkran/modules/auth"
	"github.com/nattaponra/gosongkran/modules/user"
)

func main() {
	r := gin.Default()

	app := &core.AppContext{
		Router: r,
		Config: core.AppConfig{
			JWTSecret: "secret",
			Port:      "8014",
		},
		Logger: log.Default(),
	}

	core.RegisterModules(
		&auth.AuthModule{},
		&user.UserModule{},
	)
	core.InitModules(app)

	r.Run(":" + app.Config.Port)
}
