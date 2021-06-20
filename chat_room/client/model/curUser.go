package model

import (
	"go_project/chat_room/common/message"
	"net"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
