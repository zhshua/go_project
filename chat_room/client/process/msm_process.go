package process

import (
	"encoding/json"
	"fmt"
	"go_project/chat_room/client/utils"
	"go_project/chat_room/common/message"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMsg(content string) (err error) {
	// 创建一个SmsMsg实例
	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = CurUser.UserId
	smsMsg.UserStatus = CurUser.UserStatus

	// 序列化smsMsg
	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	msg := message.Message{
		Type: message.SmsMsgType,
		Data: string(data),
	}

	// 对msg序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	// 将msg发送给服务器
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
		Buf:  make([]byte, 1024),
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg err = ", err)
	}
	return
}
