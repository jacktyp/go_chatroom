package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 1024*8)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("server 读数据长度失败",err)
		return
	}
	//buf[:4] 转成uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	n, err := conn.Read(buf[:pkgLen])
	//比对数据长度是否一致
	if n != int(pkgLen) || err != nil {
		fmt.Println("数据不一致 || 读取数据失败")
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("server - json 反序列化失败")
		return
	}
	return
	//[0 0 0 83]
	//fmt.Println("server 读取数据: ",buf[:4])
	//fmt.Println("server 读取数据:-- ",n)
}


func writePkg(conn net.Conn, data []byte) (err error) {
	//先将数据字节长度发送
	var pkg uint32
	pkg = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4],pkg)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("服务端发送数据长度-ERROR")
		return
	}

	//发送数据
	n, err = conn.Write(data)
	if n != int(pkg) || err != nil {
		fmt.Println("服务端发送数据-ERROR")
		return
	}
	return
}