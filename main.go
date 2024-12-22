package main

import (
	"github.com/text-gofr/Handler"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	Handler.InitHandlers(app)
	app.Run()
}
