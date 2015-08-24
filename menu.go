package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func createMenuHandler(w http.ResponseWriter, r *http.Request) {
	buttons := `{
		"button":[{	
			"type":"view",
			"name":"老照片",
			"url":"%s"
		},{
			"name":"咵方&聚聚",
			"sub_button":[{	
				"type":"view",
				"name":"咵方",
				"url":"%s"
			},{
				"type":"view",
				"name":"聚聚",
				"url":"%s"
			}]
		},{
			"name":"栗木人",
			"type":"view",
			"url":"%s"
		}]}
		`

	data := fmt.Sprintf(buttons, fmt.Sprintf("%s%s", gateURL, photoURL), fmt.Sprintf("%s%s", gateURL, chatURL), fmt.Sprintf("%s%s", gateURL, activityURL), fmt.Sprintf("%s%s", gateURL, contactURL))

	token := GetToken()
	res, err := http.Post(fmt.Sprintf(createMenuURL, token), "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		response(w, fmt.Sprintf("create menu err1 :%v", err))
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		response(w, fmt.Sprintf("create menu err2 :%v", err))
		return
	}

	errres, err := getResFromBody(body)
	if err != nil {
		response(w, fmt.Sprintf("create menu err3 :%v", err))
		return
	}

	response(w, fmt.Sprintf("create menu :%s", errres.ErrMsg))
}

func init() {
	handlers["createmenu"] = createMenuHandler
}
