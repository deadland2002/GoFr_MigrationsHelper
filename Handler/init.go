package Handler

import (
	"gofr.dev/pkg/gofr"
)

type RouteHandler struct {
	Path    string
	Handler func(c *gofr.Context) (interface{}, error)
	Guard   interface{}
}

type RouteInitializer struct {
	App            *gofr.App
	ParentBasePath string
}

func NewRouteInitializer(app *gofr.App, parentBasePath string) *RouteInitializer {
	return &RouteInitializer{App: app, ParentBasePath: parentBasePath}
}

func (h *RouteInitializer) Init(handler func(context RouteInitializer)) {
	handler(*h)
}

func InitHandlers(app *gofr.App) {
	var BasePath = "/"
	var Routes = NewRouteInitializer(app, BasePath)
	Routes.Init(BasicHandler)
	Routes.Init(UserHandler)
	Routes.Init(ValidatorHandler)
}
