package service

import (
	"CMS/datasource"
	"CMS/model"
	"time"

	"github.com/kataras/iris/v12"
)

func NewRent(ctx iris.Context) {
	tx := datasource.CMSdb.Begin()
	rent := new(model.Rent)
	if err := ctx.ReadJSON(&rent); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		tx.Rollback()
		return
	}
	var rentInfo datasource.Rent
	rentInfo.UserID = rent.UserID
	rentInfo.CarId = rent.CarId
	var rentStatus datasource.RentStatus
	rentStatus.CarID = rentInfo.CarId
	rentStatus.UserID = rentInfo.UserID
	rentStatus.ChargerID = 1
	rentStatus.Date = time.Now()
	datasource.CreateRentStatusPointer(tx, &rentStatus)
	rentInfo.StatusID = rentStatus.ID
	err := datasource.CreateRentPointer(tx, &rentInfo)
	if err != nil {
		data := "Exist"
		ctx.JSON(model.Response{Status: true, Data: data})
		tx.Rollback()
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
		tx.Commit()
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

func GetPayBack(ctx iris.Context) {
	type Info struct {
		UserID int `json:"userID"`
		CarID  int `json:"carID"`
	}
	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	rentInfo, err := datasource.FindRentbyUserIDCarID(datasource.CMSdb, info.UserID, info.CarID)
	if err != nil {
		data := "Not Found"
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	} else {
		type temp struct {
			Status  bool        `json:"status"`
			Data    interface{} `json:"data"`
			PayBack int         `json:"payback"`
		}
		data := "OK"
		carInfo, err := datasource.FindCarbyID(datasource.CMSdb, info.CarID)
		if err != nil {
			data := "Car Not Found"
			ctx.JSON(model.Response{Status: false, Data: data})
			return
		}
		rentStatusInfo, err := datasource.FindRentStatusbyID(datasource.CMSdb, rentInfo.StatusID)
		if err != nil {
			data := "RentStatus Not Found"
			ctx.JSON(model.Response{Status: false, Data: data})
			return
		}
		payback := carInfo.Deposite - rentStatusInfo.Fine - rentStatusInfo.Cost
		ctx.JSON(temp{Status: true, Data: data, PayBack: payback})
	}
}

func DeleteRent(ctx iris.Context) {
	tx := datasource.CMSdb.Begin()
	rent := new(model.Rent)
	if err := ctx.ReadJSON(&rent); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		tx.Rollback()
		return
	}
	err := datasource.DeleteRent(tx, rent.UserID, rent.CarId)
	if err != nil {
		data := "Not Exist"
		ctx.JSON(model.Response{Status: true, Data: data})
		tx.Rollback()
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
		tx.Commit()
	}
}
