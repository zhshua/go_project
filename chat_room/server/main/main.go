package main

import (
	"fmt"
	"net"
)

func main() {
	// 启动server程序监听套接字
	fmt.Println("server启动中...")
	Listen, err := net.Listen("tcp", "172.21.6.187:8889")
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

		// 启一个工作协程处理连接
		go func(conn net.Conn) {
			// 延时关闭连接
			defer conn.Close()
			processor := &Processor{
				Conn: conn,
			}
			err = processor.process()
			if err != nil {
				fmt.Println("processor.process() err = ", err)
				return
			}
		}(conn)

	}
}
