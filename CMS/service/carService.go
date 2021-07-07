package service

import (
	"CMS/datasource"
	"CMS/model"

	"github.com/kataras/iris/v12"
)

func NewCar(ctx iris.Context) {
	car := new(model.Car)
	if err := ctx.ReadJSON(&car); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	var carInfo datasource.Car
	carInfo.Deposite = car.Deposite
	carInfo.ID = car.ID

	err := datasource.CreateCar(datasource.CMSdb, carInfo)
	if err != nil {
		data := "Exist"
		ctx.JSON(model.Response{Status: true, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
}

// -1 for not having id and "" for not having name for search

func GetCar(ctx iris.Context) {
	type Info struct {
		ID int `json:"id"`
	}

	type temp struct {
		Status bool         `json:"status"`
		Data   interface{}  `json:"data"`
		Slice  []*model.Car `json:"slice"`
	}

	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	if info.ID != -1 {
		carInfo, err := datasource.FindCarbyID(datasource.CMSdb, info.ID)
		if err != nil {
			data := "Not Found"
			ctx.JSON((model.Response{Status: false, Data: data}))
			return
		} else {
			data := "OK"
			var car model.Car
			car.ID = carInfo.ID
			car.Deposite = carInfo.Deposite
			var slice []*model.Car
			slice = append(slice, &car)
			ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		}
	} else {
		data := "Wrong Format"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
	return
}

func GetCarInRange(ctx iris.Context) {
	type Info struct {
		Min int `json:"min"`
		Max int `json:"max"`
	}

	type temp struct {
		Status bool         `json:"status"`
		Data   interface{}  `json:"data"`
		Slice  []*model.Car `json:"slice"`
	}

	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	carsInfo, err := datasource.FindCarInRange(datasource.CMSdb, info.Min, info.Max)
	if err != nil {
		data := "Not found"
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	data := "OK"
	var slice []*model.Car
	for _, carInfo := range carsInfo {
		var car model.Car
		car.ID = carInfo.ID
		car.Deposite = carInfo.Deposite
		slice = append(slice, &car)
	}
	ctx.JSON(temp{Status: true, Data: data, Slice: slice})
}

func DeleteCar(ctx iris.Context) {
	type IDInfo struct {
		ID int `json:"ID"`
	}
	id := new(IDInfo)
	if err := ctx.ReadJSON(&id); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	err := datasource.DeleteCar(datasource.CMSdb, id.ID)
	if err != nil {
		data := "delete error"
		ctx.JSON(model.Response{Status: true, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
}
