package main

import (
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"net/http"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		//检查是否登陆
		_, err := session(writer, request)
		if err != nil {
			//没错误证明已经登陆了
			//定义一个切片，把""""内的HTML模板存到里面
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			//没登陆，走下面
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
