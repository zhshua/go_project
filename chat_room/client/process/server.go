package process

import (
	"fmt"
	"go_project/chat_room/client/utils"
	"net"
	"os"
)

// 显示登录成功的界面
func ShowMenu(){
	fmt.Println("---------恭喜xxx登录成功---------")
	fmt.Println("---------1. 显示在线用户列表---------")
	fmt.Println("---------2. 发送消息---------")
	fmt.Println("---------3. 信息列表---------")
	fmt.Println("---------4. 退出系统---------")
	fmt.Println("请选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)

	switch key{
	case 1:
		fmt.Println("显示在线用户列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("聊天历史记录")
	case 4:
		fmt.Println("你选择退出了系统")
		os.Exit(0)
	default:
		fmt.Println("输入有误，请重新输入")
	}
}

// 和服务器保持通讯
func serverProcessMsg(conn net.Conn){
	// 创建一个Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
		Buf: make([]byte, 4096),
	}
	for {
		fmt.Println("偷偷等待客户端的消息")
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err = ", err)
			return
		}
		// 等待下一步处理
		fmt.Printf("读取到消息msg = %v\n", msg)
	}
}