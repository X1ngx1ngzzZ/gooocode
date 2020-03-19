package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

//var mapUsers map[string]int
var mapUsers = make(map[string]int)

func checkerror(err error) {
	if err != nil {
		panic("error" + err.Error())
	}
}

func main() {
	//var listener net.Listener
	//var error error
	//var conn net.Conn

	fmt.Println("Starting the server ...")
	listener, error := net.Listen("tcp", "localhost:50000")
	//checkerror(error)
	if error != nil {
		fmt.Println("Error reading", error.Error())
		return //终止程序
	}

	for {
		conn, error := listener.Accept()
		//checkerror(error)
		if error != nil {
			fmt.Println("ccc")
			fmt.Println("Error reading", error.Error())
			return //终止程序
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {

	for {
		buf := make([]byte, 512)
		_, error := conn.Read(buf)
		//checkerror(error)
		if error != nil {
			fmt.Println("aaa")
			fmt.Println("Error reading", error.Error())
			return //终止程序
		}

		input := string(buf)
		//看字符串里是否含有SH子串
		//下面两个if都不走？？？？
		if strings.Contains(input, ":SH") {
			fmt.Println("server shutting down.")
			os.Exit(0)
		}

		if strings.Contains(input, ":WHO") {
			displayList()
		}
		//返回子串在字符串中第一次出现的位置
		ix := strings.Index(input, "says")
		//把says第一次出现位置之前的值用切片存储为clname
		clName := input[0 : ix-1]
		//把clname的切片类型强制转换为字符串类型存在map里，并把V置为1
		mapUsers[string(clName)] = 1
		fmt.Printf("Received: --%v--\n", string(buf))
	}
}

func displayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for k, v := range mapUsers {
		fmt.Printf("User %s is %d\n", k, v)
	}
	fmt.Println("--------------------------------------------")
}
