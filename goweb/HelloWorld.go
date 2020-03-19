package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("lisen and server:", err.Error())
	}
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello world!")

}
*/

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
