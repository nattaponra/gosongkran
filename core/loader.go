package core

import "log"

var registeredModules []Module

func RegisterModules(m ...Module) {
	registeredModules = append(registeredModules, m...)
}

func InitModules(app *AppContext) {
	for _, m := range registeredModules {
		log.Printf("Registering module: %s", m.Name())
		m.Register(app)
	}
}
