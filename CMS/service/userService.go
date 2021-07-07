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
		return
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

	type temp struct {
		Status bool          `json:"status"`
		Data   interface{}   `json:"data"`
		Slice  []*model.User `json:"slice"`
	}

	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	if info.ID != -1 && info.Name == "" {
		userInfo, err := datasource.FindUserbyID(datasource.CMSdb, info.ID)
		if err != nil {
			data := "Not Found"
			ctx.JSON((model.Response{Status: false, Data: data}))
			return
		} else {
			data := "OK"
			var user model.User
			user.ID = userInfo.ID
			user.Honesty = userInfo.Honesty
			user.IsMember = userInfo.IsMember
			user.Username = userInfo.Username
			var slice []*model.User
			slice = append(slice, &user)
			ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		}
	} else if info.ID == -1 && info.Name != "" {
		userInfos, err := datasource.FindUserbyName(datasource.CMSdb, info.Name)
		if err != nil {
			data := "Not Found"
			ctx.JSON(model.Response{Status: false, Data: data})
			return
		} else {
			data := "OK"
			var slice []*model.User
			for _, userInfo := range userInfos {
				var user model.User
				user.ID = userInfo.ID
				user.Honesty = userInfo.Honesty
				user.IsMember = userInfo.IsMember
				user.Username = userInfo.Username
				slice = append(slice, &user)
			}
			ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		}
	} else {
		data := "Wrong Format"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
	return
}

func DeleteUser(ctx iris.Context) {
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
	err := datasource.DeleteUser(datasource.CMSdb, id.ID)
	if err != nil {
		data := "delete error"
		ctx.JSON(model.Response{Status: true, Data: data})
	} else {
		data := "OK"
		ctx.JSON(model.Response{Status: true, Data: data})
	}
}

func GetUserSelf(ctx iris.Context) {
	type Info struct {
		ID  int `json:"id"`
		PWD int `json:"password"`
	}
	info := new(Info)
	if err := ctx.ReadJSON(&info); err != nil {
		ctx.StatusCode(iris.StatusOK)
		data := ""
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	if info.PWD == info.ID {
		type temp struct {
			Status bool          `json:"status"`
			Data   interface{}   `json:"data"`
			Slice  []*model.User `json:"slice"`
		}
		userInfo, err := datasource.FindUserbyID(datasource.CMSdb, info.ID)
		if err != nil {
			data := "Not Found"
			ctx.JSON((model.Response{Status: false, Data: data}))
			return
		} else {
			data := "OK"
			var user model.User
			user.ID = userInfo.ID
			user.Honesty = userInfo.Honesty
			user.IsMember = userInfo.IsMember
			user.Username = userInfo.Username
			var slice []*model.User
			slice = append(slice, &user)
			ctx.JSON(temp{Status: true, Data: data, Slice: slice})
		}
	} else {
		data := "password wrong"
		ctx.JSON(model.Response{Status: false, Data: data})
		return
	}
}
