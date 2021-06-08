package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "172.30.12.32:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err = ", err)
		} else {
			fmt.Println("成功连接一个客户端ip = ", conn.RemoteAddr().String())
		}
		go process(conn)
	}
}

func process(conn net.Conn) {

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err = ", err)
			return
		}
		fmt.Printf("收到来自客户端 %s 的消息为：%s\n", conn.RemoteAddr().String(), buf[:n])
	}
}
