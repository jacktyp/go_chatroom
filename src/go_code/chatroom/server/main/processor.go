package main

import (
	"fmt"
	"go_code/chatroom/common/message"
	userpro "go_code/chatroom/server/process"
	"go_code/chatroom/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) processAccept() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
			Buf:  [8096]byte{},
		}
		//读取数据
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("server退出")
			} else {
				fmt.Println("读取数据失败-", err)
			}
			return err
		}
		err = this.serverProcessByType(&mes)
		if err != nil {
			return err
		}
	}
}

//根据消息类型判断执行
func (this *Processor) serverProcessByType(mes *message.Message) (err error) {
	userpro := &userpro.UserProcess{
		Conn:   this.Conn,
		UserId: 0,
	}
	switch mes.Type {
	case message.LoginMesType:
		err = userpro.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//fmt.Println("register")
		err = userpro.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型不存在。。")
	}
	return
}
