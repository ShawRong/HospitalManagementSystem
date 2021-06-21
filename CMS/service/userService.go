package service

import (
	"CMS/datasource"
	"CMS/model"

	"github.com/kataras/iris/v12"
)

func NewUser(ctx iris.Context) {
	user := new(model.User)
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
	}
	var userInfo datasource.User
	userInfo.Username = user.Username
	userInfo.ID = user.ID
	userInfo.Honesty = user.Honesty
	userInfo.IsMember = user.IsMember

	err := datasource.CreateUser(datasource.CMSdb, userInfo)
	if err != nil {
		data := "Exist"
		ctx.JSON(model.Response{Status: true, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
}

// -1 for not having id and "" for not having name for search

func GetUser(ctx iris.Context) {
	type Info struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
	}
	if info.ID != -1 && info.Name == "" {

	} else if info.ID == -1 && info.Name != "" {

	} else {
		data := "Wrong Format"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
	return
}
