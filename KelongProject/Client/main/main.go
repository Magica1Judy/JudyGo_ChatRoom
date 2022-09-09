package main

import(
	"fmt"
	"LargeProject/KelongProject/Client/clientprocess"
	"os"
)

var userid string
var userpwd string
var username string


func main(){
	// 接收用户的选择
	var choose string

	for {
			fmt.Println("———————————————多人聊天系统—————————————————")
			fmt.Println("               1、登录用户")
			fmt.Println("               2、注册用户")
			fmt.Println("               3、退出系统")
			fmt.Println("——————————————————————————————————————————")
			fmt.Printf("请输入您的选择：")
			fmt.Scanln(&choose)
			switch choose{
			case "1":
				fmt.Println("1、登录用户")
				fmt.Printf("请输入您的账号：")
				fmt.Scanln(&userid)
				fmt.Printf("请输入您的密码：")
				fmt.Scanln(&userpwd)
				// 完成登录
				// 创建一个UserProcess的实例
				up := &clientprocess.UserProcess{}
				up.Login(userid, userpwd)
			case "2":
				fmt.Println("2、注册用户")
				fmt.Printf("请输入您注册账号：")
				fmt.Scanln(&userid)
				fmt.Printf("请输入您注册密码：")
				fmt.Scanln(&userpwd)
				fmt.Printf("请输入您注册名字：")
				fmt.Scanln(&username)
				//调用一个UserProcess的实例
				up := &clientprocess.UserProcess{}
				up.Register(userid, userpwd, username)
			case "3":
				fmt.Println("3、退出系统")
				//judge = false
				os.Exit(0)
			default:
				fmt.Println("输入有误，请重新输入")
			}
		}

}