package main

import (
	"flag"
	"github.com/kataras/iris/v12"
	"iris_demo/conf"
	"iris_demo/route"
)

func main() {
	flag.Parse()
	app := newApp()
	route.InitRouter(app)
	err := app.Run(iris.Addr(":"+conf.Sysconfig.Port), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}

func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	app.AllowMethods(iris.MethodOptions)
	return app
}