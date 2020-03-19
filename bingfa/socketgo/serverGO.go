package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

//把缓冲区的数据读出来并且转成字符串
/*
func read(conn net.Conn)(string,error){
	b:=make([]byte,10)
	for{
	n,error:=conn.Read(b)
	if error!=nil{
		if error==io.EOF{
			fmt.Println("Connetion is closed")
			conn.Close()
		}else{
			fmt.Println(error)
		}


	}
	content:=string(b[:n])
	return content,error
	}
}
*/

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		//设置超时时间，每个服务两秒得不到响应就关掉
		conn.SetDeadline(time.Now().Add(2 * time.Second))
		reader := bufio.NewReader(conn)
		line, error := reader.ReadString('\n')
		if error != nil {
			if error == io.EOF {
				fmt.Println("Connetion is closed")
				conn.Close()
			} else {
				fmt.Println(error)
			}
			break
		}
		fmt.Printf("recevied request %s", line)
		/*
			if(len(line)>10){
				respect:=fmt.Sprintf("dayu")

			}

		*/
	}
}

func serverGO() {

	fmt.Println("服务器开始监听")
	//返回的listen为net.listener类型，代表监听器
	listen, err := net.Listen("tcp", "127.0.0.1:8085")
	if err != nil {
		//printServerLog("Listen Error:%s",err)
		fmt.Println(err)
		return
	}
	defer listen.Close()
	//fmt.Println(listen)
	fmt.Printf("local address:%s", listen.Addr())

	// 返回的conn是一个接口类型，具有8种方法，Accept会等待客户端建立连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Established a connection with %s", conn.RemoteAddr())
		go handleConn(conn)
	}
}
/*
func clientGO() {
	conn, err := net.DialTimeout("TCP", "127.0.0.1:8085", 2*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Printf("success,remote address:%s,local address:%s", conn.RemoteAddr(), conn.LocalAddr())
	time.Sleep(2 * time.Second)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	n, error := conn.Write([]byte(line))
	if err != nil {
		fmt.Println(error)
	}
	fmt.Printf("fasong le %d", n)
}
*/
func main(){
	serverGO()
//	clientGO()
}
