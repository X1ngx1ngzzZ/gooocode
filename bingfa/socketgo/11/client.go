package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	//"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8085")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Printf("success,remote address:%s,local address:%s\n", conn.RemoteAddr(), conn.LocalAddr())
	//time.Sleep(2 * time.Second)
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("aaaaaa")
			fmt.Println(err)
		}
		n, error := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("bbbbb")
			fmt.Println(error)
		}
		fmt.Printf("fasong le %d\n", n)
	}
}
