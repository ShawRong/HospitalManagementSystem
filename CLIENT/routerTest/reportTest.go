package routerTest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ReportTest(opt string) {
	var payload *strings.Reader
	if opt == "year" {
		payload = strings.NewReader(`{` + "" + `"year":` + `true` + `}`)
	} else if opt == "quarter" {
		payload = strings.NewReader(`{` + "" + `"quarter":` + `true` + `}`)
	} else if opt == "month" {
		payload = strings.NewReader(`{` + "" + `"month":` + `true` + `}`)
	} else if opt == "day" {
		payload = strings.NewReader(`{` + "" + `"day":` + `true` + `}`)
	}
	method := "POST"
	req, err := http.NewRequest(method, RootUrl+"admin/report/report", payload)

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
