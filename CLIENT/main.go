package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"rong.com/client/routerTest"
)

func main() {
	routerTest.Init()
	/*
		routerTest.NewUserTest("Ming", true, 1)
		routerTest.GetUserTest(1, "")
		routerTest.GetUserTest(-1, "Ming")
		routerTest.DeleteUserTest(3)
	*/
	/*
		routerTest.NewCarTest(12)
		routerTest.GetCarTest(1)
		routerTest.GetCarInRangeTest(1, 2)
		routerTest.DeleteCarTest(2)
	*/
	/*
		routerTest.NewRentTest(1, 5)
		routerTest.GetRentTest(1)
		routerTest.GetPayBackTest(1, 5)
		routerTest.DeleteRentTest(1, 5)
	*/
	/*
		routerTest.GetRentStatusTest(7)
		routerTest.SetRentStatusTest(7, 21, 23, 0)
	*/
	/*
		routerTest.NewFixLogTest(1, 10)
		routerTest.GetFixLogInRangeTest(1, 20)
		routerTest.DeleteFixLogTest(10)
	*/
	/*
		routerTest.ReportTest("year")
		routerTest.ReportTest("quarter")
		routerTest.ReportTest("month")
		routerTest.ReportTest("day")
	*/
	/*
		routerTest.GetUserSelfTest(1, 1)
		routerTest.GetCarTestNormal(1)
	*/
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()

}
