package main

import (
	"fmt"
	"time"
)

func main() {
	//go println("GOGOGO")
	names := []string{"FJJ", "CQZX", "ZCY", "YT"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("name is %s\n", name)
		}(name)
		//time.Sleep(time.Second)
	}
	time.Sleep(time.Second)
}
