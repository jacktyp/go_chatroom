package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	util "go_code/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn
	//增加一个字段，表示该Conn是哪个用户
	UserId int
}

//登录
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//反序列化logindata
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("loginRes 反序列化失败")
	}

	//返回数据
	var resMes message.Message
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
		fmt.Println("返回登录数据序列化失败。。。", err)
	}
	resMes.Data = string(loginResMesData)

	resMesData, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("返回mes序列化失败")
	}
	tf := &util.Transfer{
		Conn: this.Conn,
		Buf:  [8096]byte{},
	}
	err = tf.WritePkg(resMesData)
	return
}
