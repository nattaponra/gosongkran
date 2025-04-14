package core

import "log"

var registeredModules []Module

func RegisterModule(m Module) {
	registeredModules = append(registeredModules, m)
}

func InitModules(app *AppContext) {
	for _, m := range registeredModules {
		log.Printf("Registering module: %s", m.Name())
		m.Register(app)
	}
}
