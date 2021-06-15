package main

import (
	"fmt"
	"net"
	"time"
)

func init() {
	// 初始化redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	// 初始化创建UserDao实例(需要在redisPool创建之后)
	initUserDao()
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
