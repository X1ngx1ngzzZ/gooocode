package main

import (
	"fmt"
	"time"
)

/*
func main() {
	//在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。
	//当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束
	//创建新的goroutine需要花费时间，所以会先打印world
	go hello()
	fmt.Println("world")
	time.Sleep(time.Second * 3)

}
*/
/*
func hello(i int) {
	defer mygoroutine.Done()
	fmt.Println("hello", i)

}

var mygoroutine sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		mygoroutine.Add(1)
		go hello(i)
	}
	mygoroutine.Wait()
}
*/

//channel
/*
func main() {
	//var ch1 chan int
	ch1 := make(chan int)
	//	fmt.Println(ch1)
	//无缓冲通道上发送会阻塞，所以先接收
	//主goroutine会被阻塞，所以可以开启一个goroutine来接收
	go send(10, ch1)
	go rec(ch1)
	time.Sleep(time.Second * 3)
	//	fmt.Println(s)
}

func rec(ch1 chan int) {
	<-ch1

	fmt.Println("rev,ok")

}

func send(a int, ch2 chan int) {
	ch2 <- a
	fmt.Println("send,ok")

}
*/

//for range channel

/*

func main() {
	ch1 := make(chan int)
	//	ch2 := make(chan int)
	go send(ch1)
	go rec(ch1)
	// 在主goroutine中从ch2中接收值打印
	//	for i := range ch2 { // 通道关闭后会退出for range循环
	//		fmt.Println(i)
	time.Sleep(100 * time.Second)
}

func send(ch1 chan int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
}

func rec(ch2 chan int) {
	for i := range ch2 {
		fmt.Println(i)
	}
}

*/


//单向通道channel
/*
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
*/




