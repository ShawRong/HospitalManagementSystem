package routerTest

import (
	"net/http"
)

var RootUrl string
var Client *http.Client

func Init() {

	RootUrl = "http://localhost:8080/"
	Client = &http.Client{}
}
