package service

import (
	"CMS/datasource"
	"CMS/model"
	"time"

	"github.com/kataras/iris/v12"
)

func NewRent(ctx iris.Context) {
	rent := new(model.Rent)
	if err := ctx.ReadJSON(&rent); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	var rentInfo datasource.Rent
	rentInfo.UserID = rent.UserID
	rentInfo.CarId = rent.CarId
	var err1 error
	rentInfo.Car, err1 = datasource.FindCarbyID(datasource.CMSdb, rent.CarId)
	if err1 != nil {
		ctx.StatusCode(iris.StatusOK)
		data := "No such car"
		ctx.JSON(model.Response{Status: false, Data: data})
	}
	var err2 error
	rentInfo.User, err2 = datasource.FindUserbyID(datasource.CMSdb, rent.UserID)
	if err2 != nil {
		ctx.StatusCode(iris.StatusOK)
		data := "No such user"
		ctx.JSON(model.Response{Status: false, Data: data})
	}
	var rentStatus datasource.RentStatus
	rentStatus.CarID = rentInfo.CarId
	rentStatus.UserID = rentInfo.UserID
	rentStatus.Date = time.Now()
	datasource.CreateRentStatus(datasource.CMSdb, rentStatus)
	rentInfo.StatusID = rentStatus.ID

	err := datasource.CreateRent(datasource.CMSdb, rentInfo)
	if err != nil {
		data := "Exist"
		ctx.JSON(model.Response{Status: true, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
}

func GetRent(ctx iris.Context) {
	type Info struct {
		UserID int `json:"userID"`
	}

	type temp struct {
		Status bool          `json:"status"`
		Data   interface{}   `json:"data"`
		Slice  []*model.Rent `json:"slice"`
	}

	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
	}
	if info.UserID != -1 {
		rentInfos, err := datasource.FindRentbyUserID(datasource.CMSdb, info.UserID)
		if err != nil {
			data := "Not Found"
			ctx.JSON(model.Response{Status: false, Data: data})
			return
		} else {
			data := "OK"
			var slice []*model.Rent
			for _, rentInfo := range rentInfos {
				var rent model.Rent
				rent.CarId = rentInfo.CarId
				rent.UserID = rentInfo.UserID
				rent.StatusID = rentInfo.StatusID
				slice = append(slice, &rent)
			}
			ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		}
	} else {
		data := "Wrong Format"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
	return
}
