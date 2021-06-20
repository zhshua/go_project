package process

import (
	"fmt"
	"go_project/chat_room/common/message"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

// 编写一个方法，处理返回的NotifyUserStatysMsg消息
func updateUserStatus(notifyUserStatysMsg *message.NotifyUserStatysMsg) {
	// 先判断user是否在线
	user, ok := onlineUsers[notifyUserStatysMsg.UserId]
	// user不在线，则创建一个user实例
	if !ok {
		user = &message.User{
			UserId: notifyUserStatysMsg.UserId,
		}
	}
	// user在线，直接更新状态
	user.UserStatus = notifyUserStatysMsg.Status
	onlineUsers[notifyUserStatysMsg.UserId] = user
	ShowOnlineUser()
}

// 在客户端显示当前在线的用户
func ShowOnlineUser() {
	fmt.Println("当前在线用户列表是：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}
