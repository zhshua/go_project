package main

import (
	"fmt"
	"go_project/chat_room/client/process"
	"os"
)

// 保存用户id
var userId int

// 保存用户密码
var userPwd string

func main() {
	var key int

	for true {
		fmt.Println("-------------欢迎登录多人聊天系统-------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册账号")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)

			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)

			// 创建一个UserProcess实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)

		case 2:
			fmt.Println("注册账号")

		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误请重新输入")
		}
	}
}