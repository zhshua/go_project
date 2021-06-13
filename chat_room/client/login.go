package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"go.mod/chat_room/message"
)

func login(userId int, userPwd string) (err error) {

	// 连接到服务器
	conn, err := net.Dial("tcp", "172.22.251.127:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	// 延时关闭conn连接
	defer conn.Close()

	// 实例化登录消息类型的结构体
	lgMsg := message.LoginMsg{
		UserId:  userId,
		UserPwd: userPwd,
	}
	// 对结构体序列化
	data, err := json.Marshal(lgMsg)
	if err != nil {
		fmt.Println("lgMsg json.Marshal err = ", err)
		return
	}

	// 实例化消息结构体
	msg := message.Message{
		Type: message.LoginMsgType,
		Data: string(data),
	}
	// 对结构体序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("msg json.Marshal err = ", err)
		return
	}

	var buf [4]byte // 一个uint32类型是4个字节
	// 先发送消息的长度,发送前需要用下面函数把数字类型转换成字节序, 转换结果存在buf
	binary.BigEndian.PutUint32(buf[0:4], uint32(len(data)))
	_, err = conn.Write(buf[0:4])
	if err != nil {
		fmt.Println("conn.Write head err = ", err)
		return
	}

	// 发送消息体
	_, err = conn.Write((data))
	if err != nil {
		fmt.Println("conn.Write body err = ", err)
		return
	}

	fmt.Printf("向服务器发送 %d 长度的数据, 内容是 %s\n", len(data), string(data))
	return
}
