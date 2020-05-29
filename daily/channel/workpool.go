package main

import (
	"fmt"
	"time"
)

func worker(id int,jobs <-chan int,result chan<- int){
	for j:=range jobs{
		fmt.Printf("worker:%d start job %d\n",id,j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job %d\n",id,j)
		result<- j*2
	}
}



func main(){
	//用来存工作
	jobs:=make(chan int,100)
	//
	result:=make(chan int,100)
	//开3个goroutine在等着，由于jobs里面没东西，就会被阻塞
	for w:=1;w<=3;w++{
		go worker(w,jobs,result)
	}
	//给工作
	for j:=1;j<=5;j++{
		jobs<- j
	}
	close(jobs)
	/*
	for i:=range result{
		fmt.Println(i)
	}
*/
for a := 1; a <= 5; a++ {
	k:= <-result
	fmt.Println(k)
}
}