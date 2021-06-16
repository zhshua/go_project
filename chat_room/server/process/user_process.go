/* 处理和用户相关的逻辑 */

package process

import (
	"encoding/json"
	"fmt"
	"go_project/chat_room/common/message"
	"go_project/chat_room/server/model"
	"go_project/chat_room/server/utils"
	"net"
)

// 定义处理用户例程类

type UserProcess struct {
	Conn net.Conn //连接
}

func (up *UserProcess) ServerProcessLogin(msg message.Message) (err error) {
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
	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		switch err {
		case model.ERROR_USER_PWD:
			loginResMsg.Code = 403 // 状态码403 表示密码不正确
			loginResMsg.Error = err.Error()
		case model.ERROR_USER_NOTEXIST:
			loginResMsg.Code = 500 // 状态码500 表示用户不存在
			loginResMsg.Error = err.Error()
		case model.ERROR_USER_EXIST:

		default:
			loginResMsg.Code = 505
			loginResMsg.Error = "服务器发生未知错误,请重试"
		}
	} else {
		loginResMsg.Code = 200 // 状态码200登录成功
		fmt.Printf("用户 %d 登录成功！\n", user.UserId)
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
		Conn: up.Conn,
		Buf:  make([]byte, 4096),
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("server writePkg err = ", err)
	}
	return

}

func (up *UserProcess) ServerProcessRegister(msg message.Message) (err error) {
	// 取出msg.data字段,并把它反序列化为结构体,得到login类型的消息结构体
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	// 定义回应登录消息的结构体
	var registerResMsg message.RegisterResMsg

	// 判断注册是否成功
	err = model.MyUserDao.Register(&registerMsg.User)
	if err != nil {
		if err == model.ERROR_USER_EXIST {
			registerResMsg.Code = 505
			registerResMsg.Error = model.ERROR_USER_EXIST.Error()
		} else {
			registerResMsg.Code = 506
			registerResMsg.Error = "注册时发生未知错误"
		}
	} else {
		registerResMsg.Code = 200
	}

	// 将回应类型结构体序列化
	data, err := json.Marshal(registerResMsg)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 定义消息结构体
	resMsg := message.Message{
		Type: message.RegisterResMsgType,
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
		Conn: up.Conn,
		Buf:  make([]byte, 4096),
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("server writePkg err = ", err)
	}
	return
}
