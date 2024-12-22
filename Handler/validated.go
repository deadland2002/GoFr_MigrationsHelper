package Handler

import (
	"github.com/text-gofr/Utils"
	migrations "github.com/text-gofr/migrations"
	"github.com/text-gofr/models/validator_models"
)

func ValidatorHandler(context RouteInitializer) {
	context.App.Migrate(migrations.All())
	var BasePath = "validator"

	var user = RouteHandler{
		Path:    Utils.JoinPaths(context.ParentBasePath, BasePath, "user"),
		Handler: validator_models.GetUser,
	}

	context.App.GET(user.Path, user.Handler)
}
