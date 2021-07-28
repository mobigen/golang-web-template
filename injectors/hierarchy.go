package injectors

import (
	"github.com/jblim0125/golang-web-platform/common"
	"github.com/jblim0125/golang-web-platform/infrastructures/datastore"
	"github.com/jblim0125/golang-web-platform/infrastructures/router"
)

// Injector web-server layer initializer : Dependency Injection )
type Injector struct {
	Router    *router.Router
	Datastore *datastore.DataStore
	Log       *common.Logger
}

// New create Injector
func (Injector) New(r *router.Router, d *datastore.DataStore,
	l *common.Logger) *Injector {
	return &Injector{
		Router:    r,
		Datastore: d,
		Log:       l,
	}
}

// Init init web-server layer interconnection create (web server layer init
func (h *Injector) Init() error {
	// path grouping
	apiv1 := h.Router.Group("/api/v1")

	// Todo
	h.Log.Errorf("[ PATH ] /api/v1/todos ............................................................ [ OK ]")
	todo := Todo{}.Init(h)
	apiv1.GET("/todos", todo.GetAll)
	apiv1.GET("/todos/:id", todo.GetByID)
	apiv1.POST("/todos", todo.Create)

	// Ex : Controller - Application - Domain
	// user := User{}.Init(initializer.Datastore)
	// apiv1.GET("/todos", todo.GetAll)
	// apiv1.GET("/todos/:id", todo.GetByID)
	// apiv1.POST("/todos", todo.Create)

	h.Log.Errorf("[ PATH ] /api/v1/ping ............................................................. [ OK ]")
	ping := Ping{}.Init(h)
	apiv1.GET("/ping", ping.GetPing)

	return nil
}
