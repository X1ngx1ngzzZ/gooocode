package main

import (
	"fmt"
	"net"
	//"time"
	//"bufio"
	//	"io"
)

func main() {

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
	//监听并接收连接
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

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		//设置超时时间，每个服务两秒得不到响应就关掉
		//别特么瞎定时间！！！！
		//conn.SetDeadline(time.Now().Add(2 * time.Second))
		/*
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
		*/
		buf := make([]byte, 512)
		lens, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:lens]))
		//fmt.Printf("recevied request %s", line)
		/*
			buff:=make([]byte,512)
				if(len(lens)>10){
				respect:=fmt.Sprintf("dayu")

				}

		*/
		lenss, err := conn.Write(buf)
		fmt.Println("fasongchenggong1", string(buf[:lenss]))

	}
}
