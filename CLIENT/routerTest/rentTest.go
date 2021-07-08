package routerTest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func NewRentTest(userID int, carID int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"userID":` + strconv.Itoa(userID) + `,` + `"carID":` + strconv.Itoa(carID) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/rent/new", payload)

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

func GetRentTest(userID int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"userID":` + strconv.Itoa(userID) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/rent/get", payload)

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

func GetPayBackTest(userID int, carID int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"userID":` + strconv.Itoa(userID) + `,` + `"carID":` + strconv.Itoa(carID) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/rent/getPayBack", payload)

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

func DeleteRentTest(userID int, carID int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"userID":` + strconv.Itoa(userID) + `,` + `"carID":` + strconv.Itoa(carID) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/rent/delete", payload)

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
