package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL)

	fmt.Fprintln(w, "-----------------------")

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))

	fmt.Fprintln(w, "-----------------------")
	r.ParseForm()
	fmt.Fprintln(w, r.Form)

	fmt.Fprintln(w, "-----------------------")

	fmt.Fprintln(w, r.Header)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
