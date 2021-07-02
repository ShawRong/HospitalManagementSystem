package controller

import (
	"CMS/service"

	"github.com/kataras/iris/v12"
)

func HubController(app *iris.Application) {
	/*router main*/
	main := app.Party("/")

	/*router home*/
	//home := main.Party("/")

	/*router user*/
	user := main.Party("/user")
	user.Post("/get", service.GetUser)
	user.Post("/new", service.NewUser)
	user.Post("/delete", service.DeleteUser)

	/*router car*/
	car := main.Party("/car")
	car.Post("/get", service.GetCar)
	car.Post("/new", service.NewCar)
	car.Post("/delete", service.DeleteCar)

	/*router rent*/
	rent := main.Party("/rent")
	rent.Post("/get", service.GetRent)
	rent.Post("/new", service.NewRent)
}
