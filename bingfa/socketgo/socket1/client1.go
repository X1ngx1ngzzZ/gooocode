package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, error := net.Dial("tcp", "localhost:50000")
	//checkError(error)
	if error!= nil {
		fmt.Println("Error reading", error.Error())
		return //终止程序
	}

	//新建一个缓冲区，把键盘输入的暂存在里面
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	//最后的\n应该是给socket让它可以断句
	clientname, _ := inputReader.ReadString('\n')
	//去掉输入的里面的字符串里的\r\n
	trimmedClient := strings.Trim(clientname, "\r\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit. Type SH to shutdown server.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, error = conn.Write([]byte(trimmedClient + "says" + trimmedInput))
		//checkError(error)
		if error!= nil {
			fmt.Println("aaa")
			fmt.Println("Error reading", error.Error())
			return //终止程序
		}

	}

}

func checkError(err error) {
	if err != nil {
		panic("Error: " + err.Error()) // terminate program
	}
}
