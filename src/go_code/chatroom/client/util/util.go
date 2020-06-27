package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

//传输消息
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //这时传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("server 读数据长度失败", err)
		return
	}
	//buf[:4] 转成uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	//比对数据长度是否一致
	if n != int(pkgLen) || err != nil {
		fmt.Println("数据不一致 || 读取数据失败")
		return
	}

	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("server - json 反序列化失败")
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先将数据字节长度发送
	var pkg uint32
	pkg = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[:4], pkg)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("服务端发送数据长度-ERROR")
		return
	}
	//发送数据
	n, err = this.Conn.Write(data)
	if n != int(pkg) || err != nil {
		fmt.Println("服务端发送数据-ERROR")
		return
	}
	return
}
