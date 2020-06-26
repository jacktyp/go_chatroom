package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server 8889 listen")
	//本机-8889端口
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listen err=: ", err)
		return
	}
	//一直监听
	for {
		fmt.Println("等待客户端连接服务器。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err=: ", err)
		}
		//连接成功
		//启动协程
		go process(conn)

	}
}

func process(conn net.Conn) {
	defer conn.Close()
	pro := &Processor{Conn: conn}
	err := pro.processAccept()
	if err != nil {
		fmt.Println("客户端-服务器通信协程--error")
		return
	}
}
