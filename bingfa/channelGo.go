package main

import (
	"fmt"
	"time"
)
/*
func main() {
	ch := make(chan string)

	go send(ch)
	go get(ch)
	time.Sleep(2 * time.Second)
}

func send(ch chan string) {
	ch <- "wo"
	ch <- "shi"
	ch <- "ni"
	ch <- "die"
}

func get(ch chan string) {
	var input string
	//time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s", input)
	}
}
*/
/*
func main() {
	ch1 := make(chan int)
	go pump(ch1) // pump hangs
	go suck(ch1)
	v, ok := <-ch1
	fmt.Println(v, ok) // prints only 0
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
*/
/*
func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	//如果先传进去，没人接，传阻塞，取得时候里面没有，取阻塞
	//所以要先取后传
	//out <- 2
	go f1(out)
	out <- 2
}
*/



//实现锁与信号量
/*
type Empty interface{}
type semaphore chan Empty

func (s semaphore)P(n int){
	e:=new(Empty)
	for i:=0;i<n;i++{
		s<-e
	}
}

func (s semaphore)V(n int){
	for i:=0;i<n;i++{
		<-s
	}
}

 /* mutexes *
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock(){
	s.V(1)
}

// signal-wait 
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}
*/

func main() {
	stream := pump()
	go suck(stream)
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}








