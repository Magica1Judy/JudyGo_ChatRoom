package clientprocess

import(
	"fmt"
	"LargeProject/KelongProject/Common/message"
	"encoding/json"
	"LargeProject/KelongProject/Common/control"
	//"sync"
	
	
)

type MesProcess struct{

}


func (this *MesProcess)TalkToSomeBody()(){

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>您成功进入私聊, 您的发言只会被目标人接收")
	fmt.Printf("请输入您想要私聊的用户id:")
	var TalkToId string
	fmt.Scanln(&TalkToId)
	
	_, ok := ShowClientOnlineUsers[TalkToId]
	if !ok{
		fmt.Println("抱歉，您想要私聊的用户不在线")
		return
	}
	this.SendToSomeBodyMes(TalkToId)

	
	
}

func (this *MesProcess)SendToSomeBodyMes(talktoid string)(){
	//var mtx sync.Mutex
	var TalkText string
	
	var mes message.Message
	mes.Type = message.TalkToSomeBodyType
	
	servicer := &message.SercetTalk{
		Content : make(chan string, 5),
		UserId : talktoid,
	}
	defer close(servicer.Content)

	

	for TalkText != "exit"{
		fmt.Printf("请输入你的发言(输入 exit 为退出): ")
		fmt.Scanln(&TalkText)
		if TalkText == "exit"{
			return
		}
		//mtx.Lock()
		ReadySendMsg := "["+LoginedUser.UserId+"]号用户"+"对你说:"+TalkText
		servicer.Content <-ReadySendMsg
		//mtx.Unlock()
	// 	fmt.Printf(ReadySendMsg)
	// }

	// for {
		//mtx.Lock()
		msg := <- servicer.Content
		//mtx.Unlock()
		var talkdata message.TalkData
		talkdata.Content = msg
		talkdata.UserId = talktoid


		data, err := json.Marshal(talkdata)
		if err != nil{
			fmt.Println("小序列化talkdata错误, 原因是：", err.Error())
			return
		}

		mes.Data = string(data)

		data, err = json.Marshal(mes)
		if err != nil{
			fmt.Println("大序列化talkdata错误, 原因是：", err.Error())
			return
		}

		tf := &control.Transfer{
			Conn : LoginedUser.Conn,
		}
		err = tf.WritePkg(data)
		if err != nil{
			fmt.Println("发送mes错误, 原因是：", err.Error())
			return
		}
	}
}





func (this *MesProcess)FreeTalkAbout()(){

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>您成功进入公共聊天室, 可以进行任意发言")
	var TalkText string
	for TalkText != "exit"{
		fmt.Printf("请输入你的发言(输入 exit 为退出): ")

		fmt.Scanln(&TalkText)
		this.SendGroupMes(TalkText)
	}
}


func (this *MesProcess)SendGroupMes(content string)(){

	if content == "exit"{
		return
	}
	// 还是以message形式进行传输
	var mes message.Message
	mes.Type = message.SendFreeMesType

	var talkdata message.TalkData
	talkdata.Content = content
	talkdata.UserId = LoginedUser.UserId
	talkdata.UserStatus = LoginedUser.UserStatus

	data, err := json.Marshal(talkdata)
	if err != nil{
		fmt.Println("小序列化talkdata错误, 原因是：", err.Error())
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil{
		fmt.Println("大序列化talkdata错误, 原因是：", err.Error())
		return
	}

	tf := &control.Transfer{
		Conn : LoginedUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("发送mes错误, 原因是：", err.Error())
		return
	}
}
