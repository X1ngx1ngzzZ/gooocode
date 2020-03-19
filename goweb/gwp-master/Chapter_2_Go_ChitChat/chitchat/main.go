package main

import (
	"net/http"
	"time"
)

func main() {
	p("ChitChat", version(), "started at", config.Address)

	// handle static assets
	//多路复用器
	mux := http.NewServeMux()
	//Filesever创建静态文件处理器，用mux.handle调用，把config.static的数据作为files
	files := http.FileServer(http.Dir(config.Static))
	//当传过来以/static为开头的URL时候，就移除/static（StripPrefix的作用）去files里面找
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", index)
	// error
	mux.HandleFunc("/err", err)

	// defined in route_auth.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	//对用户的身份进行验证，并在验证成功后返回一个cookie
	mux.HandleFunc("/authenticate", authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
