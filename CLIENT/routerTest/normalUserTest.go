package routerTest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetUserSelfTest(id int, password int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"id":` + strconv.Itoa(id) + `,` + "" + `"password":` + strconv.Itoa(password) + `}`)
	req, err := http.NewRequest(method, RootUrl+"normal/getUserSelf", payload)

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

func GetCarTestNormal(id int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"id":` + strconv.Itoa(id) + `}`)
	req, err := http.NewRequest(method, RootUrl+"normal/getCar", payload)

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
