package main

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"io"
	"net"
)

func main() {
	fmt.Println("server 8889 listen")
	//本机-8889端口
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listen err=: ",err)
		return
	}
	//一直监听
	for {
		fmt.Println("等待客户端连接服务器。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err=: ",err)
		}
		//连接成功
		//启动协程
		go process(conn)
	}
}

//根据消息类型判断执行
func serverProcessByType(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
		case message.LoginMesType :
			err = serverProcessLogin(conn,mes)
		case message.RegisterMesType:
			fmt.Println("register")
		default:
			fmt.Println("消息类型不存在。。")
	}
	return
}

//登录
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//反序列化logindata
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("loginRes 反序列化失败")
	}

	//返回数据
	var  resMes message.Message
	//类型
	resMes.Type = message.LoginResMesType
	var loginResMes message.LoginResMes

	//登录校验
	if loginMes.UserId == 159 && loginMes.UserPwd == "123456" {
		fmt.Println("ID，密码-校验合法。。")
		loginResMes.Code = 200
	} else {
		fmt.Println("不合法。。。")
		loginResMes.Code = 500
		loginResMes.Error = "用户不存在，请先注册。。。"
	}
	//序列化-封装
	loginResMesData, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("返回登录数据序列化失败。。。",err)
	}
	resMes.Data = string(loginResMesData)

	resMesData, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("返回mes序列化失败")
	}

	err = writePkg(conn,resMesData)
	return
}


func process(conn net.Conn) {
	defer conn.Close()

	for {
		//读取数据
		mes,err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("server退出")
				return
			}
			fmt.Println("读取数据失败-",err)
			return
		}
		//fmt.Println(mes)
		err = serverProcessByType(conn,&mes)
		if err != nil {
			return
		}

	}
}

