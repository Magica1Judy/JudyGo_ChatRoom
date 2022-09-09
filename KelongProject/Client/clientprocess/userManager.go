package clientprocess

import(
	"LargeProject/KelongProject/Client/model"
	"LargeProject/KelongProject/Common/message"
	"fmt"
	"encoding/json"
	"LargeProject/KelongProject/Common/control"
)

var ShowClientOnlineUsers map[string]*message.User = make(map[string]*message.User, 10)

//首先完成对 单个 用户变量的初始化
var LoginedUser model.CurUser


// 对服务端发送的数据进行处理 发送的是一个结构体变量 id 和 用户状态
func UpdateOnlineUser(notifyUserStatusMes *message.NotifyUserStatusMes){

	//进行判断 看是否在 ShowClientOnlineUsers 该切片中已经存在了上线id
	user, ok := ShowClientOnlineUsers[notifyUserStatusMes.UserId]
	// 如果不存在
	if !ok {
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}
	}
	// userstatus 是message.User里面的参数
	user.UserStatus = notifyUserStatusMes.Status
	ShowClientOnlineUsers[notifyUserStatusMes.UserId] = user
	// 对它们进行输出
	// OutputShowClientOnlineUsers()
}

func OutputShowClientOnlineUsers(){
	fmt.Println("    【当前在线用户列表如下：】")
	for id, _ := range ShowClientOnlineUsers{
		fmt.Println("用户id : ", id)
	}
}


func Offline(){
	delete(ShowClientOnlineUsers,LoginedUser.UserId)
	// 还是以message形式进行传输
	var mes message.Message
	mes.Type = message.ExitFreeRoomType

	var offlineid message.OfflineData
	offlineid.UserId = LoginedUser.UserId
	

	data, err := json.Marshal(offlineid)
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
