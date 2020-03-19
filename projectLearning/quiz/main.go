package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//注册flag，
	//三个参数：名称，默认值，提示信息
	csvFilename := flag.String("csv", "problems.csv", "a csv file in bulabulabula")
	//注册flag
	flag.Parse()
	//_ = csvFilename
	file, err := os.Open(*csvFilename)
	if err != nil {
		/*
		fmt.Printf("failed to open csv:%s", *csvFilename)
		os.Exit(1)
		*/
		//使用退出函数替代
		exit(fmt.SPrintf("failed to open csv:%s", *csvFilename))
	}
	_ = file
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}