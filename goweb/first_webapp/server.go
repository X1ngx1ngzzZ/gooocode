package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

//http.Request 描述了客户端请求，内含一个 URL 字段。
/*

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

*/

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!!!!!!!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called-!!!!"+name)
		h(w,r)
	}
}
func main() {
	server:=http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello",log(hello))
	server.ListenAndServe()
}
