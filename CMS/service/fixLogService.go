package service

import (
	"CMS/datasource"
	"CMS/model"
	"time"

	"github.com/kataras/iris/v12"
)

func NewFixLog(ctx iris.Context) {
	type Info struct {
		CarID int
		Fee   int
	}
	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	var fixLogInfo datasource.FixLog
	fixLogInfo.CarID = info.CarID
	fixLogInfo.Fee = info.Fee
	fixLogInfo.Time = time.Now()
	err := datasource.CreateFixLog(datasource.CMSdb, fixLogInfo)
	ctx.StatusCode(iris.StatusOK)
	if err != nil {
		data := "Error"
		ctx.JSON(model.Response{Status: false, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
}

func GetFixLogInRange(ctx iris.Context) {
	type Info struct {
		Min int `json:"min"`
		Max int `json:"max"`
	}

	type temp struct {
		Status bool            `json:"status"`
		Data   interface{}     `json:"data"`
		Slice  []*model.FixLog `json:"slice"`
	}

	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	fixLogsInfo, err := datasource.FindFixLogInRange(datasource.CMSdb, info.Min, info.Max)
	if err != nil {
		data := "Not found"
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	data := "OK"
	var slice []*model.FixLog
	for _, fixLogInfo := range fixLogsInfo {
		var fixLog model.FixLog
		fixLog.ID = fixLogInfo.ID
		fixLog.CarID = fixLogInfo.CarID
		fixLog.Fee = fixLogInfo.Fee
		fixLog.Time = fixLogInfo.Time
		slice = append(slice, &fixLog)
	}
	ctx.JSON(temp{Status: true, Data: data, Slice: slice})
}

func DeleteFixLog(ctx iris.Context) {
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
	err := datasource.DeleteFixLog(datasource.CMSdb, id.ID)
	if err != nil {
		data := "delete error"
		ctx.JSON(model.Response{Status: false, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
}
