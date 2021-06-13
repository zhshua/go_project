package process

import (
	"chat_room/message"
	"chat_room/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

// 定义处理用户例程类
type UserProcess struct {
	Conn net.Conn //连接
}

func (this *UserProcess) ServerProcessLogin(msg message.Message) (err error) {
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

	// 创建一个发送数据的实例去发送数据
	tf := &utils.Transfer{
		Conn: this.Conn,
		Buf:  make([]byte, 4096),
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("server writePkg err = ", err)
	}
	return

}
