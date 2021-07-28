package injectors

import (
	"github.com/jblim0125/golang-web-platform/controllers"
)

// Ping ...
type Ping struct{}

// Init for interconnection [ controller(App) - Service(Repository) - repository - datastore ]
func (Ping) Init(i *Injector) *controllers.Ping {
	return controllers.Ping{}.New()
}
