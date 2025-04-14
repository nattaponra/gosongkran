package core

import "gorm.io/gorm"

type Module interface {
	Name() string
	Register(app *AppContext)
}

type Migratable interface {
	Migrate(db *gorm.DB) error
}
