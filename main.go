package main

import (
	"github.com/kataras/iris/v12"
	"gses2-btc-application/api"
)

var mainPagePath = "./gses2.app"
var mainPageRoute = "/gses2.app" // must be the same as the mainPagePath
var apiRoute = mainPageRoute + "/api"

// Middleware is a middleware that logs the path of the request.
func Middleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path: %s", ctx.Path())
	ctx.Next()
}

func main() {
	app := iris.New()
	app.Use(Middleware)

	api.ConnectTo(app, apiRoute)

	app.Get("/", func(ctx iris.Context) {
		ctx.Redirect(mainPageRoute)
	})

	app.HandleDir(mainPageRoute, mainPagePath, iris.DirOptions{
		IndexName: "index.html",
	})

	_ = app.Run(iris.Addr(":8080"))
}
