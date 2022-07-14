package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

/**
author: Wenjie_pan
*/
var urlUse string
var userName string
var password string
var times int

func getRequestInfo() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.Method == "GET" {
			if len(r.Form["url"]) != 0 && len(r.Form["userName"]) != 0 && len(r.Form["password"]) != 0 {
				times++
				urlUse = r.Form["url"][0]
				fmt.Print(urlUse)
				userName = r.Form["userName"][0]
				password = r.Form["password"][0]
			}
		}
	})
	http.ListenAndServe(":9000", nil)
}

func postInfo() string {
	dataAll := make(url.Values)
	dataAll.Add("userName", userName)
	dataAll.Add("passowrd", password)
	payload := dataAll.Encode()
	bodyget, _ := http.Post(
		urlUse,
		"application/x-www-form-urlencoded",
		strings.NewReader(payload),
	)
	content, _ := ioutil.ReadAll(bodyget.Body)
	f, _ := os.Create("first.txt")
	getVal := string(content)
	io.WriteString(f, getVal)
	return string(content)

}

func running() {
	for {
		for {
			if times < 1 {
				break
			}
			fmt.Print(0)
			postInfo()
			time.Sleep(time.Second)
			times--
		}
	}

}

func main() {
	go running()
	getRequestInfo()
}
