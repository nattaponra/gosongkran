package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nattaponra/gosongkran/core"
)

func main() {
	r := gin.Default()

	app := &core.AppContext{
		Router: r,
		Config: map[string]interface{}{},
		Logger: log.Default(),
	}

	core.InitModules(app)

	r.Run(":8080")
}
