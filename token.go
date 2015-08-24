package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	accessToken []byte
)

type SuccessRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetToken() string {
	return string(accessToken)
}

func getAccessToken() ([]byte, error) {
	res, err := http.Get(fmt.Sprintf(tokenURL, APPID, APPSECRET))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	errres, err := getResFromBody(body)
	if err != nil {
		return nil, err
	}

	if errres != nil {
		return nil, errors.New(errres.ErrMsg)
	} else {
		var result SuccessRes
		err = json.Unmarshal(body, &result)
		if err != nil {
			return nil, err
		}

		return []byte(result.AccessToken), nil
	}
}

func init() {
	go func() {
		for {
			var err error
			accessToken, err = getAccessToken()
			if err != nil {
				fmt.Println("get access token err is ", err)
				continue
			}
			time.Sleep(time.Hour * 1)
		}
	}()
}
