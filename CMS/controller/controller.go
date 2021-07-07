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

	/*rounter nomarl user*/
	normalUser := main.Party("/normal")
	normalUser.Post("/getCar", service.GetCar)
	normalUser.Post("/getUserSelf", service.GetUserSelf)
	/*rounter admin*/
	admin := main.Party("/admin")

	/*router user*/
	user := admin.Party("/user")
	user.Post("/get", service.GetUser)
	user.Post("/new", service.NewUser)
	user.Post("/delete", service.DeleteUser)

	/*router car*/
	car := admin.Party("/car")
	car.Post("/get", service.GetCar)
	car.Post("/new", service.NewCar)
	car.Post("/delete", service.DeleteCar)
	car.Post("/getInRange", service.GetCarInRange)

	/*router rent*/
	rent := admin.Party("/rent")
	rent.Post("/get", service.GetRent)
	rent.Post("/new", service.NewRent)
	rent.Post("/getPayBack", service.GetPayBack)

	/*router rentStatus*/
	rentStatus := admin.Party("/rentStatus")
	rentStatus.Post("/get", service.GetRentStatus)
	rentStatus.Post("/set", service.SetRentStatus)

	/*router fixLog*/
	fixLog := admin.Party("/fixLog")
	fixLog.Post("/new", service.NewFixLog)
	fixLog.Post("/getInRange", service.GetFixLogInRange)
	fixLog.Post("/delete", service.DeleteFixLog)

	/*router report*/
	report := admin.Party("/report")
	report.Post("/report", service.NewReport)
}
