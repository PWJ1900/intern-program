package main

import (
	"net/http"
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
				userName = r.Form["userName"][0]
				password = r.Form["password"][0]
				_, res, _ := LoginBind(userName, password, urlUse)
				w.Write([]byte(res))
			}
		}
	})
	http.ListenAndServe(":9000", nil)
}

func main() {
	getRequestInfo()
}
