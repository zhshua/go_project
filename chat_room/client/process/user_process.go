package process

import (
	"encoding/json"
	"fmt"
	"go_project/chat_room/client/utils"
	"go_project/chat_room/message"
	"net"
)

// 定义UserProcess类
type UserProcess struct{

}

func (up *UserProcess)Login(userId int, userPwd string) (err error) {

	// 连接到服务器
	conn, err := net.Dial("tcp", "172.21.6.187:8889")
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

	// 创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
		Buf: make([]byte, 4096),
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("login writePkg err = ", err)
		return
	}
	fmt.Printf("向主机 %s 发送长度为 %d 数据, 内容是 %s\n", conn.RemoteAddr().String(), len(data), string(data))

	msg, err = tf.ReadPkg()
	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	if loginResMsg.Code == 200 {
		// fmt.Println("登录成功")
		go serverProcessMsg(conn)

		// 显示登录成功后的二级菜单
		for{
			ShowMenu()
		}
	} else {
		fmt.Println("登录失败, err = ", loginResMsg.Error)
	}
	return
}