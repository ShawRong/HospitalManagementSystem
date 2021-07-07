package service

import (
	"CMS/datasource"
	"CMS/model"

	"github.com/kataras/iris/v12"
)

func GetRentStatus(ctx iris.Context) {
	type Info struct {
		ID int `json:"ID"`
	}

	type temp struct {
		Status bool                `json:"status"`
		Data   interface{}         `json:"data"`
		Slice  []*model.RentStatus `json:"slice"`
	}

	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
	}
	if info.ID != -1 {
		rentStatusInfo, err := datasource.FindRentStatusbyID(datasource.CMSdb, info.ID)
		if err != nil {
			data := "Not Found"
			ctx.JSON(model.Response{Status: false, Data: data})
			return
		} else {
			data := "OK"
			var rentStatus model.RentStatus
			rentStatus.CarID = rentStatusInfo.CarID
			rentStatus.ChargerID = rentStatusInfo.ChargerID
			rentStatus.Cost = rentStatusInfo.Cost
			rentStatus.Date = rentStatusInfo.Date
			rentStatus.Fine = rentStatusInfo.Fine
			rentStatus.ID = rentStatusInfo.ID
			rentStatus.UserID = rentStatusInfo.UserID
			var slice []*model.RentStatus
			slice = append(slice, &rentStatus)
			ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		}
	} else {
		data := "Wrong Format"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
	return
}

func SetRentStatus(ctx iris.Context) {
	type Info struct {
		ID        int `json:"ID"`
		Fine      int `json:"fine"`
		Cost      int `json:"cost"`
		ChargerID int `json:"chargerID"`
	}
	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
	}
	rentStatusInfo, err := datasource.FindRentStatusbyID(datasource.CMSdb, info.ID)
	if err != nil {
		data := "Not Found"
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	} else {
		rentStatusInfo.Cost = info.Cost
		rentStatusInfo.Fine = info.Fine
		rentStatusInfo.ChargerID = info.ChargerID
	}
	err = datasource.UpdateRentStatus(datasource.CMSdb, info.ID, rentStatusInfo)
	if err != nil {
		data := "Set Error"
		ctx.JSON(model.Response{Status: false, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
}
