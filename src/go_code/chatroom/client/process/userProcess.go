package process

import (
	"encoding/json"
	"fmt"
	util "go_code/chatroom/client/util"
	common "go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
	"os"
)

type UserProcess struct {
}

func (userProcess *UserProcess) Login(userId int, userPwd string) (err error) {
	fmt.Println(userId, userPwd)
	//连接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("client 连接失败")
		return err
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
	if err != nil {
		fmt.Println("loginData转换json 失败")
		return err
	}
	mes.Data = string(loginData)

	//将mes转json
	mesData, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes转换json 失败")
		return err
	}
	tf := &util.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	tf.WritePkg(mesData)
	fmt.Println("client send message!--")

	//接受服务器消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("接收服务器消息失败")
		return err
	}
	var loginResMes common.LoginResMes
	json.Unmarshal([]byte(mes.Data), &loginResMes)

	//登录成功
	if loginResMes.Code == 200 {
		//显示在线用户列表
		fmt.Println("在线用户列表")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户ID：\t", v)
		}

		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯.如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端.
		go serverProcessMes(conn)
		//显示菜单
		for {
			ShowMenu()
		}
		//fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	fmt.Println(mes)
	return
}

//注册
func (this *UserProcess) Register(userId int,
	userPwd string, userName string) (err error) {

	//链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2. 准备通过conn发送消息给服务
	var mes common.Message
	mes.Type = common.RegisterMesType
	//3. 创建一个LoginMes 结构体
	var registerMes common.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将registerMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 5. 把data赋给 mes.Data字段
	mes.Data = string(data)

	// 6. 将 mes进行序列化化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//创建一个Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	mes, err = tf.ReadPkg() // mes 就是 RegisterResMes

	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成 RegisterResMes
	var registerResMes common.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功, 你重新登录一把")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
