package injectors

import (
	"github.com/mobigen/golang-web-template/common"
	"github.com/mobigen/golang-web-template/infrastructures/datastore"
	"github.com/mobigen/golang-web-template/infrastructures/router"
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
	// For Version
	ver := Version{}.Init(h)
	h.Router.GET("/version", ver.GetVersion)

	// path grouping
	apiv1 := h.Router.Group("/api/v1")

	// Sample
	h.Log.Errorf("[ PATH ] /api/v1/sample ........................................................... [ OK ]")
	sample := Sample{}.Init(h)
	apiv1.GET("/samples", sample.GetAll)
	apiv1.GET("/sample/:id", sample.GetByID)
	apiv1.POST("/sample", sample.Create)
	apiv1.POST("/sample/update", sample.Update)
	apiv1.DELETE("/sample/:id", sample.Delete)

	return nil
}
