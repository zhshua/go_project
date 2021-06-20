/* 主控器, 去分发各类处理逻辑 */

package main

import (
	"errors"
	"fmt"
	"go_project/chat_room/common/message"
	"go_project/chat_room/server/process"
	"go_project/chat_room/server/utils"
	"io"
	"net"
)

// 定义processor类
type Processor struct {
	Conn net.Conn
}

// 分发处理消息类型
func (p *Processor) serverProcessMsg(msg message.Message) (err error) {

	// 看看是否能收到客户端发来的群发消息
	fmt.Println("msg = ", msg)

	switch msg.Type {
	case message.LoginMsgType:
		// 处理登录消息
		// 创建一个UserProcess实例
		up := &process.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(msg)
	case message.RegisterMsgType:
		// 处理注册消息
		// 创建一个UserProcess实例
		up := &process.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessRegister(msg)
	case message.SmsMsgType:
		smsProcess := process.SmsProcess{}
		smsProcess.SendGroupMsg(&msg)
	default:
		err = errors.New("不支持的消息类型")
		return
	}
	return
}

func (p *Processor) process() (err error) {
	for {
		// 创建传输数据实例
		tf := &utils.Transfer{
			Conn: p.Conn,
			Buf:  make([]byte, 4096),
		}

		var msg message.Message
		msg, err = tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端已经关闭连接了, 服务器端也关闭连接退出")
				return
			} else {
				fmt.Println("readPkg err = ", err)
				return
			}
		}
		fmt.Printf("读取到来自 %s 地址发来的数据为:\n%v\n", p.Conn.RemoteAddr().String(), msg)

		err = p.serverProcessMsg(msg)
		if err != nil {
			fmt.Println("serverProcessMsg err = ", err)
			return
		}
	}
}
