package core

type Module interface {
	Name() string
	Register(app *AppContext)
}
