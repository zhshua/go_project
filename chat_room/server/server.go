package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"go.mod/chat_room/message"
)

func readPkg(conn net.Conn) (msg message.Message, err error) {
	// 1. conn.Read返回的是读取到的字节数
	// 2. conn,Read会把读取的[]byte结果存储到传进的buf切片中
	buf := make([]byte, 4096)
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read head err = ", err)
		return
	}

	// 把字节序转换成uint32类型的数字
	pkgLen := binary.BigEndian.Uint32(buf[:4])
	_, err = conn.Read(buf[:pkgLen])
	if err != nil {
		fmt.Println("conn.Read body err = ", err)
		return
	}

	// 反序列化结构体
	json.Unmarshal(buf[:pkgLen], &msg)
	return
}

func process(conn net.Conn) {
	// 延时关闭连接
	defer conn.Close()
	for {
		msg, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端已经关闭连接了, 服务器端也关闭连接退出")
				return
			} else {
				fmt.Println("readPkg err = ", err)
				return
			}
		}
		fmt.Printf("读取到来自 %s 地址发来的数据为:\n%v\n", conn.RemoteAddr().String(), msg)
	}
}

func main() {
	// 启动server程序监听套接字
	fmt.Println("server启动中...")
	Listen, err := net.Listen("tcp", "172.22.251.127:8889")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	for {
		conn, err := Listen.Accept()
		if err != nil {
			fmt.Println("Listen.Accept err = ", err)
			continue
		}
		fmt.Printf("已经连接到ip地址为 %s 的客户端\n", conn.RemoteAddr().String())
		go process(conn)

	}
}
