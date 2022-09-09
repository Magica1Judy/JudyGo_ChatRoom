package model

import(
	"net"
	"LargeProject/KelongProject/Common/message"
)

// 单个用户变量？？？？作为全局变量可以使代码更加简洁？？？
type CurUser struct{
	Conn net.Conn
	message.User
}