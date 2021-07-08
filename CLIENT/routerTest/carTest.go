package routerTest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func NewCarTest(deposite int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"deposite":` + strconv.Itoa(deposite) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/car/new", payload)

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

func GetCarTest(id int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"id":` + strconv.Itoa(id) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/car/get", payload)

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

func GetCarInRangeTest(min int, max int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"min":` + strconv.Itoa(min) + `,` + `"max":` + strconv.Itoa(max) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/car/getInRange", payload)

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

func DeleteCarTest(id int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"id":` + strconv.Itoa(id) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/car/delete", payload)

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
