package message

// 定义几个消息类型常量
const (
	LoginMsgType    = "LoginMsg"
	LoginResMsgType = "LoginResMsg"
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

// 登录的消息类型
type LoginMsg struct {
	UserId  int    `json:"userId"`  // 登录用户id
	UserPwd string `json:"userPwd"` // 登录用户密码
}

// 接收登录信息的消息类型
type LoginResMsg struct {
	Code  int    // 错误码
	Error string // 错误信息
}
