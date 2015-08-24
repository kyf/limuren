package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	if err != nil {
		response(w, fmt.Sprintf("error:%v", err))
		return
	}

	var rev Text
	err = xml.Unmarshal(body, &rev)
	if err != nil {
		response(w, fmt.Sprintf("error:%v", err))
		return
	}

	var res TextReply = TextReply{}
	res.ToUserName = rev.FromUserName
	res.FromUserName = rev.ToUserName
	res.CreateTime = fmt.Sprintf("%v", time.Now().Unix())
	res.Content = fmt.Sprintf("receive data is %s", rev.Content)
	res.MsgType = "text"

	re, err := xml.MarshalIndent(res, "", "\t")
	if err != nil {
		response(w, fmt.Sprintf("convert data xml error : %v", err))
		return
	}

	sre := strings.Replace(string(re), "TextReply", "xml", -1)

	response(w, sre)
}

func validHandler(w http.ResponseWriter, r *http.Request) {
	params := getForm(r)
	signature := params.Get("signature")
	timestamp := params.Get("timestamp")
	nonce := params.Get("nonce")
	echostr := params.Get("echostr")

	tmp := [...]string{TOKEN, timestamp, nonce}
	s := sort.StringSlice(tmp[0:])
	sort.Sort(s)
	v := sha1.Sum([]byte(strings.Join(s, "")))
	if strings.EqualFold(hex.EncodeToString(v[:]), signature) {
		w.Write([]byte(echostr))
	} else {
		w.Write([]byte("error"))
	}
}
