package clientprocess

import(
	"fmt"
	"os"
	"net"
	"LargeProject/KelongProject/Common/control"
	"LargeProject/KelongProject/Common/message"
	"encoding/json"
)

func ShowMenu(){
	for{
		var key string
		//var content string
		// 因为我们总会SmsProcess实例，因此我们将其定义在switch外部
		mesProcess := &MesProcess{}

		fmt.Println("——————————菜单——————————")
		fmt.Println("    1、显示在线用户列表")
		fmt.Println("    2、发送消息")
		fmt.Println("    3、信息列表")
		fmt.Println("    4、（测试）请勿打扰 先不接受群聊信息")
		fmt.Println("    5、私聊用户")
		fmt.Println("    6、退出系统")
		fmt.Println("————————————————————————")
		fmt.Printf("请输入您的选项：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			OutputShowClientOnlineUsers()
		case "2":
			mesProcess.FreeTalkAbout()
			// fmt.Printf("请输入你的发言: ")
			// fmt.Scanln(&content)
			// mesProcess.SendGroupMes(content)
		case "3":
			fmt.Println("    3、信息列表")
		case "4":
			Offline()
		case "5":
			mesProcess.TalkToSomeBody()
		case "6":
			os.Exit(0)
		default:
			fmt.Println("输入不正确")
	}
	}
}

// 保持和服务器端通信
func serverProcessMes(conn net.Conn)(){
	//需要一个transfer实例， 不停读取服务器发送的信息
	tf := &control.Transfer{
		Conn : conn,
	}
	fmt.Println("********客户端在不断读取服务器发送的数据********")
	for{
		//只要对方不关闭连接， 就一直堵塞
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("服务器端断开链接:",err)
			return
		}

		//读到服务端发来的消息 下一步处理
		switch mes.Type {

		//若为关于在线用户列表的消息种类
		case message.NotifyUserStatusMesType:
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			UpdateOnlineUser(&notifyUserStatusMes)
			
		case message.SendFreeMesType:
			outputMestoROOM(&mes)

		case message.TalkToSomeBodyType:
			outputMestoSomeBody(&mes)




		}
	}

}
