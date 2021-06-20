package message

// 定义几个消息类型常量
const (
	LoginMsgType            = "LoginMsg"
	LoginResMsgType         = "LoginResMsg"
	RegisterMsgType         = "RegisterMsg"
	RegisterResMsgType      = "RegisterResMsg"
	NotifyUserStatysMsgType = "NotifyUserStatysMsg"
)

// 定义几个User状态
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息内容
}

// "登录"的消息类型
type LoginMsg struct {
	UserId   int    `json:"userId"`   // 登录用户id
	UserPwd  string `json:"userPwd"`  // 登录用户密码
	UserName string `json:"userName"` // 登录的用户名
}

// "回应登录信息"的消息类型
type LoginResMsg struct {
	Code    int    `json:"code"`  // 状态码
	Error   string `json:"error"` // 错误信息
	UsersId []int  // 存放所有在线用户的id
}

// "注册"消息的类型
type RegisterMsg struct {
	User User `json:"user"` // User结构体
}

// "回应注册信息"的消息类型
type RegisterResMsg struct {
	Code  int    `json:"code"`  // 状态码
	Error string `json:"error"` // 错误信息
}

// "服务器主动推送通知用户状态"的消息类型
type NotifyUserStatysMsg struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}
