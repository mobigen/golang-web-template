package injectors

import (
	"github.com/jblim0125/golang-web-platform/controllers"
	"github.com/jblim0125/golang-web-platform/repositories"
	"github.com/jblim0125/golang-web-platform/services"
)

// Todo ...
type Todo struct{}

// Init for interconnection [ controller(App) - Service(Repository) - repository - datastore ] : Dependency Injection
func (Todo) Init(in *Injector) *controllers.Todo {
	repository := repositories.Todo{}.New(in.Datastore)
	service := services.Todo{}.New(repository)
	return controllers.Todo{}.New(service)
}
