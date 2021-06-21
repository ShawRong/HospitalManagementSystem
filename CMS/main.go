package main

import (
	"CMS/controller"
	"CMS/datasource"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	controller.HubController(app)
	err := datasource.DBConnect()
	if err != nil {
		println(err)
	}
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
	defer datasource.DBClose(datasource.CMSdb)
}
