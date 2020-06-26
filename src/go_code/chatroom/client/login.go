package main

import (
	"encoding/json"
	"fmt"
	common "go_code/chatroom/common/message"
	"net"
)

func login(userId int,userPwd string) {
	fmt.Println(userId,userPwd)
	//连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("client 连接失败")
	}
	defer conn.Close()
	//封装数据
	var mes common.Message
	mes.Type = common.LoginMesType
	//封装登录数据
	var loginMes common.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//转换成json
	loginData, err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("loginData转换json 失败")
		return
	}
	mes.Data = string(loginData)

	//将mes转json
	mesData, err := json.Marshal(mes)
	if err != nil{
		fmt.Println("mes转换json 失败")
		return
	}

	writePkg(conn,mesData)
	fmt.Println("client send message!--")

	//接受服务器消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("接收服务器消息失败")
		return
	}
	var loginResMes common.LoginResMes
	json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	fmt.Println(mes)

	return
}