package routerTest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func NewFixLogTest(carID int, Fee int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"carID":` + strconv.Itoa(carID) + `,` + `"fee":` + strconv.Itoa(Fee) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/fixLog/new", payload)

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

func GetFixLogInRangeTest(min int, max int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"min":` + strconv.Itoa(min) + `,` + `"max":` + strconv.Itoa(max) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/fixLog/getInRange", payload)

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

func DeleteFixLogTest(id int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"ID":` + strconv.Itoa(id) + `}`)
	fmt.Println(payload)
	req, err := http.NewRequest(method, RootUrl+"admin/fixLog/delete", payload)

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
