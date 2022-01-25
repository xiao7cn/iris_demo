package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris_demo/controller"
)

func InitRouter(app *iris.Application) {
	bathUrl := "/api"
	mvc.New(app.Party(bathUrl +"/user")).Handle(controller.NewUserController())
}