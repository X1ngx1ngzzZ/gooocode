package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//URL
	//fmt.Fprintln(w, r.URL)
	//form
	//r.ParseForm()
	//fmt.Fprintln(w, r.URL)
	//fmt.Fprintln(w, r.Form["hello"])
	//fmt.Fprintln(w, r.Form)
	//fmt.Fprintln(w, "fenge")
	//body
	/*
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	*/
	//r.ParseMultipartForm(1024)
	//fmt.Fprintln(w,r.MultipartForm)
	//从requeset里面读取1024字节的数据
	r.ParseMultipartForm(1024)
	//取出文件头
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
			fmt.Fprintln(w, "-----------------")
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()

}
