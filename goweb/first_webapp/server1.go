package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    " 127.0.0.1:8085",
		Handler: nil,
	}

	server.ListenAndServe()
	//server.ListenAndServeTLS("cert.pem", "key.pem")
	server.ListenAndServeTLS("cert.pem", "key.pem")
}

/*
type Server struct {
　　Addr　　　　　 string
　　Handler　　　　Handler
　　ReadTimeout　　time.Duration
　　WriteTimeout　 time.Duration
　　MaxHeaderBytes int
　　TLSConfig　　　*tls.Config
　　TLSNextProto　 map[string]func(*Server, *tls.Conn, Handler)
　　ConnState　　　func(net.Conn, ConnState)
　　ErrorLog　　　 *log.Logger
}
*/
