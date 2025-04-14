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

	config := core.LoadConfig()

	config.DBDriver = "postgres"
	config.DBHost = "localhost"
	config.DBPort = "5432"
	config.DBUser = "postgres"
	config.DBPass = "postgres"
	config.DBName = "postgres"

	app := &core.AppContext{
		Router: r,
		Config: config,
		Logger: log.Default(),
		DB:     core.InitDatabase(config),
	}

	core.RegisterModules(
		&auth.AuthModule{},
		&user.UserModule{},
	)
	core.InitModules(app)

	r.Run(":" + app.Config.Port)
}
