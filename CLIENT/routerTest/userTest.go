package routerTest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func NewUserTest(username string, isMemeber bool, honesty int) {
	method := "POST"
	var temp string
	if isMemeber {
		temp = "true"
	} else {
		temp = "false"
	}
	payload := strings.NewReader(`{` + "" + `"username":` + `"` + username + `"` + `,` + "" + `"isMember":` + temp + `,` + "" + `"honesty":` + strconv.Itoa(honesty) + "" + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/user/new", payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func GetUserTest(id int, username string) {
	if id <= 0 {
		id = -1
	}
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"id":` + strconv.Itoa(id) + `,` + "" + `"name":` + `"` + username + `"` + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/user/get", payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func DeleteUserTest(id int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"id":` + strconv.Itoa(id) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/user/delete", payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := Client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
