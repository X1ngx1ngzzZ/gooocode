package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,world!")
}

type daring struct{}

func (h *daring) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,baby!")
}

func fjj(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,fjj!")
}

func main() {
	//多个处理器的实例化
	world := MyHandler{}
	baby := daring{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		//将默认处理器设置为handler
		//Handler: &handler,
	}
	//使用Handler函数进行处理器绑定
	http.Handle("/hello", &world)
	http.Handle("/baby", &baby)
	//Handlefunc可以把函数a转换成一个带方法a的handler
	http.HandleFunc("/fjj", fjj)
	server.ListenAndServe()
}
//没弄懂还
//把处理器函数转化为处理器的办法
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	//默认多路复用器调用HandleFunc方法传入URL和处理器函数
	　　 DefaultServeMux.HandleFunc(pattern, handler)
	}
	
//	而下面是ServeMux.HandleFunc 方法的定义：
//	
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter,
	　　 *Request)) {
	　　mux.Handle(pattern, HandlerFunc(handler))
	}
	