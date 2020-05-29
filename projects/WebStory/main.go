package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"./cyoa"
)

func main() {

	filename := flag.String("file", "gopher.json", "the JSON file with story")
	port := flag.Int("port", 8080, "the port to start CYOA Web Application")
	flag.Parse()
	fmt.Printf("using story in %s.\n", *filename)

	//打开文件
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	//把json文件进行解析，放在map里面返回回来
	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", story)

	h := cyoa.NewHandler(story)
	fmt.Printf("Started server at Port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
