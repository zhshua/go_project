package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
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
	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度给对方
	var buf [4]byte // 一个uint32类型是4个字节
	pkgLen := uint32(len(data))
	// 先发送消息的长度,发送前需要用下面函数把数字类型转换成字节序, 转换结果存在buf
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write head err = ", err)
		return
	}

	// 发送消息体
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write body err = ", err)
		return
	}
	return
}

func serverProcessLogin(conn net.Conn, msg message.Message) (err error) {
	// 取出msg.data字段,并把它反序列化为结构体,得到login类型的消息结构体
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	// 定义回应登录消息的结构体
	var loginResMsg message.LoginResMsg
	// 判断登录是否成功
	if loginMsg.UserId == 123 && loginMsg.UserPwd == "abcde" {
		loginResMsg.Code = 200 // 状态码200登录成功
	} else {
		loginResMsg.Code = 500 // 状态码登录失败
		loginResMsg.Error = "用户尚未注册, 请注册后再登录"
	}

	// 将回应类型结构体序列化
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 定义消息结构体
	resMsg := message.Message{
		Type: message.LoginResMsgType,
		Data: string(data),
	}
	// 将消息结构体序列化
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}
	// 发送数据
	err = writePkg(conn, data)
	if err != nil {
		fmt.Println("server writePkg err = ", err)
	}
	return

}

// 分发处理消息类型
func serverProcessMsg(conn net.Conn, msg message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType:
		// 处理登录消息
		err = serverProcessLogin(conn, msg)
	case message.RegisterMsgType:
		// 处理注册消息
	default:
		err = errors.New("不支持的消息类型")
		return
	}
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

		err = serverProcessMsg(conn, msg)
		if err != nil {
			fmt.Println("serverProcessMsg err = ", err)
			return
		}
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
