package core

var registeredModules []Module

func RegisterModules(m ...Module) {
	registeredModules = append(registeredModules, m...)
}

func InitModules(app *AppContext) {

	for _, m := range registeredModules {
		app.Logger.Printf("Registering module: %s", m.Name())
		m.Register(app)

		// Optional migration
		if mg, ok := m.(Migratable); ok {
			if err := mg.Migrate(app.DB); err != nil {
				app.Logger.Printf("Migration failed for module %s: %v", m.Name(), err)
			}
		}
	}
}
