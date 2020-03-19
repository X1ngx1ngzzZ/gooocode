package main

import "fmt"

import "time"



//channelbase1,空结构体不占内存，所以建议发送信号的都用结构体
func main() {
	chan0 := make(chan string, 3)
	chan1 := make(chan struct{}, 3)
	chan2 := make(chan struct{}, 3)

	//接收
	go func() {
		//把chan1里面的输出来，目前里面没有所以chan1阻塞
		<-chan1
		fmt.Println("Received  chan1 and wait second[receiver]")
		//输出chan0的
		for {
			if elem, ok := <-chan0; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("stopped[receiver]")
		//给chan2发信息，由于没人读所以阻塞，为撒要两个{}
		chan2 <- struct{}{}
	}()

	//发送
	go func() {
		//循环输出到chan0中
		for _, elem := range []string{"a", "b", "c", "d"} {
			chan0 <- elem
			fmt.Println("sent", elem, "[sender]")
			//当输出C时，发个信号给chan1，让接收的协程开始运行
			if elem == "c" {
				chan1 <- struct{}{}
				fmt.Println("sent a signal,[sender]")
			}
		}
		fmt.Println("wait 2 second")
		time.Sleep(2 * time.Second)
		close(chan0)
		//阻塞chan2
		chan2 <- struct{}{}
	}()
	//chan2的存在就是为了让main协程停下来等chan0和chan1
	<-chan2
	<-chan2
}


//selectuse

func main(){
	
}