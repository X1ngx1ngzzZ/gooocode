package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f,_:=os.Create("trace.out")
	defer f.Close()
	err:=trace.Start(f)
	if err!=nil{
		panic(err)
	}
	defer trace.Stop()
	//s := []int{0, 1, 2, 3, 4}
	fmt.Println("hello world")
}
