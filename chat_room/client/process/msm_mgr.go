package process

import (
	"encoding/json"
	"fmt"
	"go_project/chat_room/common/message"
)

func OutputGroupMsg(msg *message.Message) {
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	fmt.Printf("用户id:\t%d对大家说:\t%s\n\n", smsMsg.UserId, smsMsg.Content)

}
