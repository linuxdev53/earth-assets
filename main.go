package main

import (
	"earth-assets/common"
	"earth-assets/controllers"

	"github.com/kataras/iris/v12"
)

// main function
func main() {
	// create app
	app := iris.New()
	app.HandleDir("/", "./assets", iris.DirOptions{
		IndexName: "/index.html",
		Gzip:      false,
		ShowList:  false,
	})

	// endpoints
	app.Get("/api/EarthAssets", controllers.EarthAssets)

	app.Run(iris.Addr(common.ListenAddr))
}
