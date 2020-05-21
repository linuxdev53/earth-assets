package main

import (
	"earth-assets/common"
	"earth-assets/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// main function
func main() {
	// create app
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Load the template files.
	tmpl := iris.HTML("./views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)

	app.HandleDir("/public", "./views/public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		ctx.View("shared/error.html")
	})

	// endpoints
	api := mvc.New(app.Party("/api/EarthAssets"))
	api.Handle(new(controllers.EarthAssets))

	app.Run(
		iris.Addr(common.ListenAddr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
