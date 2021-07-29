package injectors

import (
	"github.com/mobigen/golang-web-template/controllers"
)

// Version version injector
type Version struct{}

// Init version controller create
func (Version) Init(in *Injector) *controllers.Version {
	return controllers.Version{}.New()
}
