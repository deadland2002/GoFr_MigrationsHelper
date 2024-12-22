package Handler

import (
	"github.com/text-gofr/Utils"
	"gofr.dev/pkg/gofr"
)

func greet(c *gofr.Context) (interface{}, error) {
	return "HI", nil
}

func BasicHandler(context RouteInitializer) {
	var BasePath = "basic"

	var greetRoute = RouteHandler{
		Path:    "greet",
		Handler: greet,
	}

	context.App.GET(Utils.JoinPaths(context.ParentBasePath, BasePath, greetRoute.Path), greetRoute.Handler)
}
