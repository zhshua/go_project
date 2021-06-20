/* 处理和短消息相关的逻辑 */
package process

import (
	"encoding/json"
	"fmt"
	"go_project/chat_room/common/message"
	"go_project/chat_room/server/utils"
	"net"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMsg(msg *message.Message) {
	// 遍历服务端的onlinesUsers将消息取出并转发出去

	// 反序列化SmsMsg
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		// 防止自己给自己转发
		if id == smsMsg.UserId {
			continue
		}
		sp.SendMsgToEachOnlineUser(data, up.Conn)
	}
}

func (sp *SmsProcess) SendMsgToEachOnlineUser(data []byte, conn net.Conn) {
	// 创建一个Transfer 实例去发送消息
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  make([]byte, 1024),
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("消息转发失败~")
	}
}
