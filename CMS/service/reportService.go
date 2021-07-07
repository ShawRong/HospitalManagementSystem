package service

import (
	"CMS/datasource"
	"CMS/model"

	"github.com/kataras/iris/v12"
)

type Item struct {
	Earn int `json:"earn"`
	Time int `json:"time"`
}

func NewReport(ctx iris.Context) {
	type Info struct {
		Year    bool `json:"year"`
		Quarter bool `josn:"quarter"`
		Month   bool `json:"month"`
		Day     bool `json:"day"`
	}
	type temp struct {
		Status bool        `json:"status"`
		Data   interface{} `json:"data"`
		Slice  []*Item     `json:"slice"`
	}
	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	if info.Year {
		items, err := datasource.GetReportViewYear(datasource.CMSdb)
		ctx.StatusCode(iris.StatusOK)
		if err != nil {
			data := "error"
			ctx.JSON((model.Response{Status: false, Data: data}))
			return
		}
		data := "OK"
		var slice []*Item
		for _, itemInfo := range items {
			var item Item
			item.Earn = itemInfo.Earn
			item.Time = itemInfo.Time
			slice = append(slice, &item)
		}
		ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		return
	} else if info.Quarter {
		items, err := datasource.GetReportViewQuarter(datasource.CMSdb)
		ctx.StatusCode(iris.StatusOK)
		if err != nil {
			data := "error"
			ctx.JSON((model.Response{Status: false, Data: data}))
			return
		}
		data := "OK"
		var slice []*Item
		for _, itemInfo := range items {
			var item Item
			item.Earn = itemInfo.Earn
			item.Time = itemInfo.Time
			slice = append(slice, &item)
		}
		ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		return
	} else if info.Month {
		items, err := datasource.GetReportViewMonth(datasource.CMSdb)
		ctx.StatusCode(iris.StatusOK)
		if err != nil {
			data := "error"
			ctx.JSON((model.Response{Status: false, Data: data}))
			return
		}
		data := "OK"
		var slice []*Item
		for _, itemInfo := range items {
			var item Item
			item.Earn = itemInfo.Earn
			item.Time = itemInfo.Time
			slice = append(slice, &item)
		}
		ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		return
	} else if info.Day {
		items, err := datasource.GetReportViewDay(datasource.CMSdb)
		ctx.StatusCode(iris.StatusOK)
		if err != nil {
			data := "error"
			ctx.JSON((model.Response{Status: false, Data: data}))
			return
		}
		data := "OK"
		var slice []*Item
		for _, itemInfo := range items {
			var item Item
			item.Earn = itemInfo.Earn
			item.Time = itemInfo.Time
			slice = append(slice, &item)
		}
		ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		return
	}

}
