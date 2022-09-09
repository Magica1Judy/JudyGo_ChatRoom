package serverprocess

import(
	"fmt"
	"LargeProject/KelongProject/Common/message"
	"encoding/json"
	"LargeProject/KelongProject/Common/control"
	"net"
)

type ServerProcess struct{

}

// 处理发过来的消息
func (this *ServerProcess) ServerSendGroupMes(mes *message.Message){
	var talkdata message.TalkData

	// 先对发来的数据进行解码
	err := json.Unmarshal([]byte(mes.Data), &talkdata)
	if err != nil {
		fmt.Println("在服务端对客户端发来的信息Unmarshal失败，原因是: ",err)
		return
	}

	// 然后进行对不同用户进行转发
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("在服务端对客户端发来的信息marshal失败，原因是: ",err)
		return
	}
	for id, up := range userManager.onlineUsers{
		// 不要再发给自己
		if id == talkdata.UserId{
			continue
		}
		this.SendMesToEachUser(data, up.Conn)
	}
}


func (this *ServerProcess) SendMesToEachUser(data []byte, conn net.Conn){
	tf := &control.Transfer{
		Conn : conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败，原因是: ",err)
	}
}


// 处理发过来的消息
func (this *ServerProcess) ServerSendSomeBodyMes(mes *message.Message){
	var talkdata message.TalkData

	// 先对发来的数据进行解码
	err := json.Unmarshal([]byte(mes.Data), &talkdata)
	if err != nil {
		fmt.Println("在服务端对客户端发来的信息Unmarshal失败，原因是: ",err)
		return
	}

	// 然后进行对目标用户进行转发
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("在服务端对客户端发来的信息marshal失败，原因是: ",err)
		return
	}
	

	up, err := userManager.GetOnlineUserById(talkdata.UserId)
	// for id, up := range userManager.onlineUsers{
	// 	// 发给目标用户
	// 	if id == talkdata.UserId{
			
	// 	}else{
	// 		fmt.Println("server经过查找目标用户不在线或者不存在")
	// 		return
	// 	}
	// }
	if err != nil{
		fmt.Println("server经过查找目标用户不在线或者不存在")
		return
	}else{
		this.SendMesToSomeBody(data, up.Conn)
	}
}

func (this *ServerProcess) SendMesToSomeBody(data []byte, conn net.Conn){
	tf := &control.Transfer{
		Conn : conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToSomeBody 转发消息失败，原因是: ",err)
	}
}