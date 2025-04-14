# ⚙️ GoSongkran – Plug-and-Play Go API Framework
GoSongkran is a lightweight, plug-and-play API framework for Golang, designed to help you build scalable, modular backends with ease.

It enables fast development of RESTful services by letting you register modules (like Auth, User, Payments, etc.) independently. Each module is self-contained and auto-registers itself into the main app – no spaghetti wiring.


## 🚀 Features
- ✅ Modular Architecture – Just import and use.
- 🧩 Plug-and-Play Modules – Auto-register on boot.
- 🛡️ Built-in JWT Auth Support
- 📦 Centralized Context – Share Router, Logger, Config, DB.
- 🧪 Testable – Handlers and services can be easily tested.
- 📄 Ready for Extensions – Add middlewares, plugins, Swagger, etc.


```
gosongkran/
│
├── core/                   # Framework core (context, loader, interfaces)
│   ├── context.go
│   ├── module.go
│   └── loader.go
│
├── modules/                # Your API modules
│   ├── auth/
│   │   ├── module.go       # Self-registered module
│   │   ├── handler.go
│   │   └── service.go
│   └── user/
│       └── ...
│
└── main.go                 # Entry point

```


## 🧠 Concept
1. 🔌 Plug & Register Modules
Each module implements the core.Module interface and auto-registers itself using init():

go
```go
type Module interface {
	Name() string
	Register(app *AppContext)
}
```

2. 🔁 Self-Register on Import
```go
func init() {
	core.RegisterModule(&AuthModule{})
}
```
Just import the module once, and it’s active!

3. 🔧 AppContext Injection
Every module receives shared resources via AppContext:
```go
type AppContext struct {
	Router *gin.Engine
	Logger *log.Logger
	Config map[string]interface{}
}
```


### ✨ Example Usage
main.go
```go
import (
	"github.com/nattaponra/gosongkran/core"
	_ "github.com/nattaponra/gosongkran/modules/auth"
	_ "github.com/nattaponra/gosongkran/modules/user"
)

func main() {
	app := &core.AppContext{
		Router: gin.Default(),
		Logger: log.Default(),
	}

	core.InitModules(app)
	app.Router.Run(":8080")
}
```

### 🔐 Auth Module Example
```go
func (m *AuthModule) Register(app *core.AppContext) {
	group := app.Router.Group("/auth")
	group.POST("/login", LoginHandler)
}
```

# 📌 Status
✅ MVP ready
🚧 Coming soon:

- DB/ORM integration
- Swagger module loader
- CLI scaffold tool
- gRPC support

# 📜 License
MIT