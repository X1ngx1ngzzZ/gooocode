package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatData(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	//创建FuncMap的映射，把映射的K设为函数名字，V为实际定义的函数
	funcMap := template.FuncMap{"fdate": formatData}
	//将FuncMap与模板进行绑定
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("tmpl.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()

}
