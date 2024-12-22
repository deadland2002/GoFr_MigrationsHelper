package Handler

import (
	"github.com/text-gofr/Utils"
	"github.com/text-gofr/models/user_models"
)

func UserHandler(context RouteInitializer) {
	var BasePath = "user"

	var userGetAll = RouteHandler{
		Path:    "getAll",
		Handler: user_models.GetUserList,
	}

	var userAdd = RouteHandler{
		Path:    "create",
		Handler: user_models.AddNewUser,
	}

	var userEdit = RouteHandler{
		Path:    "edit",
		Handler: user_models.EditUser,
	}

	context.App.GET(Utils.JoinPaths(context.ParentBasePath, BasePath, userGetAll.Path), userGetAll.Handler)
	context.App.POST(Utils.JoinPaths(context.ParentBasePath, BasePath, userAdd.Path), userAdd.Handler)
	context.App.PATCH(Utils.JoinPaths(context.ParentBasePath, BasePath, userEdit.Path), userEdit.Handler)
}
