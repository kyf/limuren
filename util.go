package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type ErrorRes struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func getResFromBody(body []byte) (*ErrorRes, error) {
	var tmp map[string]interface{}
	err := json.Unmarshal(body, &tmp)
	if err != nil {
		return nil, err
	}

	if _, ok := tmp["errcode"]; ok {
		var res ErrorRes
		err = json.Unmarshal(body, &res)
		if err != nil {
			return nil, err
		}

		return &res, nil
	} else {
		return nil, nil
	}
}

func response(w http.ResponseWriter, content interface{}) {
	w.Header().Add("Content-Type", "application/xml")
	switch c := content.(type) {
	case []byte:
		w.Write(c)
	case string:
		w.Write([]byte(c))
	default:
		re, err := xml.MarshalIndent(c, "", "\t")
		if err == nil {
			w.Write(re)
		}
	}
}
