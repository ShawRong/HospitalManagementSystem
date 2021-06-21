package controller

import (
	"github.com/kataras/iris/v12"
)

func HubController(app *iris.Application) {
	/*router main*/
	main := app.Party("/")

	/*router home*/
	home := main.Party("/")
	home.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})
}
