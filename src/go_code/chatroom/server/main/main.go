package main

import (
	"fmt"
	"go_code/chatroom/server/model"
	"net"
	"time"
)

//初始化
func init() {
	//当服务器启动时，我们就去初始化我们的redis的连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

//完成对UserDao的初始化任务
func initUserDao() {
	// pool 是一个全局的变量
	model.MyUserDao = model.NewUserDao(pool)
}
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
