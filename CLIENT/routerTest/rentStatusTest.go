package routerTest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetRentStatusTest(id int) {
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"id":` + strconv.Itoa(id) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/rentStatus/get", payload)

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

func SetRentStatusTest(id int, fine int, cost int, chargerID int) {
	if chargerID <= 0 {
		chargerID = 1
	}
	method := "POST"
	payload := strings.NewReader(`{` + "" + `"id":` + strconv.Itoa(id) + `,` + `"fine":` + strconv.Itoa(fine) + `,` + `"cost":` + strconv.Itoa(cost) + `,` + `"chargerID":` + strconv.Itoa(chargerID) + `}`)
	req, err := http.NewRequest(method, RootUrl+"admin/rentStatus/set", payload)

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
