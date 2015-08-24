package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type handler struct{}

type processor func(w http.ResponseWriter, r *http.Request)

var (
	handlers map[string]processor = make(map[string]processor)
)

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Replace(r.URL.Path, "/", "", 1)
	if p, ok := handlers[path]; ok {
		p(w, r)
		return
	}

	if strings.EqualFold(path, "") {
		params := getForm(r)
		echostr := params.Get("echostr")

		if !strings.EqualFold(echostr, "") {
			validHandler(w, r)
		} else {
			messageHandler(w, r)
		}
	} else {
		http.NotFound(w, r)
	}
}

const (
	TOKEN     string = "kyf"
	HTTP_PORT string = ":1472"
	APPID     string = "wxc2fbad3afaa5adb1"
	APPSECRET string = "f131a927fc806372823dac2ae6b81d26"
)

func getForm(r *http.Request) url.Values {
	r.ParseForm()
	r.ParseMultipartForm(1024 * 1024 * 1024)
	return r.Form
}

func main() {
	fmt.Println("service is running ...")
	err := http.ListenAndServe(HTTP_PORT, &handler{})
	if err != nil {
		fmt.Println("web server error is ", err)
	}
}
