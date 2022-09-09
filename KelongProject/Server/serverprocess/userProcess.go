package serverprocess

import(
	"net"
	"LargeProject/KelongProject/Common/message"
	"encoding/json"
	"fmt"
	"LargeProject/KelongProject/Common/control"
	"LargeProject/KelongProject/Server/model"

)


type UserProcess struct{
	Conn net.Conn
	// 增加一个字段 表示该UserProcess是属于哪个用户的
	Process_userid string
}


func (this *UserProcess)OfflineUsers(mes *message.Message){
	var offlineData message.OfflineData
	err := json.Unmarshal([]byte(mes.Data), &offlineData)
	if err != nil {
		fmt.Println("在ServerProcess的OfflineUsers信息反序列化失败, 原因是",err)
		return
	}
	userManager.DelOnlineUser(offlineData.UserId)
}


func (this *UserProcess)ServerProcessLogin(mes *message.Message)(err error){
	// 先从mes中取出它的mes.Dat, 并直接反序列化为LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("在ServerProcess的userProcess登录信息反序列化失败, 原因是",err)
		return
	}

	var ReturnMes message.Message
	ReturnMes.Type = message.LoginReturnType

	// 另一个结构体 里面要装错误信息 ErrorNum ErrorData
	var loginReturnMes message.LoginReturnMes
	_, err = model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS{
			loginReturnMes.ErrorNum = 400 // 400不合法，用户不存在
			loginReturnMes.ErrorData = err.Error()
		}else if err == model.ERROR_USER_PWD{
			loginReturnMes.ErrorNum = 401 // 密码不正确
			loginReturnMes.ErrorData = err.Error()
		}else{
			loginReturnMes.ErrorNum = 505
			loginReturnMes.ErrorData = "服务器内部错误。。"
		}	
	}else{
		loginReturnMes.ErrorNum = 200 // 200合法
		//登录成功的id
		this.Process_userid = loginMes.UserId
		
		//具有全局变量的Manager可以进行直接调用
		//调用其增添方法
		userManager.AddOrUpdateOnlineUser(this)
		this.NotifyOtherOnlineUsers(loginMes.UserId)


		for id, _ := range userManager.onlineUsers{
			loginReturnMes.AllLoginUsers = append(loginReturnMes.AllLoginUsers, id)
		}

		
		fmt.Println("登录成功")
	}




	// 将 loginReturnMes序列化
	data, err := json.Marshal(loginReturnMes)
	if err != nil {
		fmt.Println("在UserProcess中登录检验结果序列化失败, 原因是",err)
		return
	}
	ReturnMes.Data = string(data)
	data, err = json.Marshal(ReturnMes)
	if err != nil {
		fmt.Println("在UserProcess中登录检验结果序列化失败, 原因是",err)
		return
	}
	tf := &control.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (this *UserProcess)ServerProcessRegister(mes *message.Message)(err error){
	var Registermes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &Registermes)
	if err != nil {
		fmt.Println("在ServerProcess中获取注册信息反序列化失败, 原因是",err)
		return
	}

	var returnmes message.Message
	returnmes.Type = message.RegisterReturnType
	
	var returnRegistermes message.RegisterReturnMes
	err = model.MyUserDao.Register(&Registermes)
	if err != nil{
		if err == model.ERROR_USER_EXISTS{
			returnRegistermes.ErrorNum = 505 // 505不合法，用户已经存在
			returnRegistermes.ErrorData = err.Error()
		}else{
			returnRegistermes.ErrorNum = 506 
			returnRegistermes.ErrorData = "注册发生未知错误"
		}
	}else{
		returnRegistermes.ErrorNum = 200
	}

	data, err := json.Marshal(returnRegistermes)
	if err != nil {
		fmt.Println("serverProcess的returnRegistermes注册检验结果序列化失败, 原因是",err)
		return
	}

	returnmes.Data = string(data)
	data, err = json.Marshal(returnmes)
	if err != nil {
		fmt.Println("登录检验结果序列化失败, 原因是",err)
		return
	}
	// 发送data数据 也需要丢包检测 为了减少代码冗杂 故也封装函数writePkg
	// 因为使用了分层模式 先创建 transfer实例
	tf := &control.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return


}






func (this *UserProcess)NotifyOtherOnlineUsers(userid string){
	// 首先遍历现存登录了的用户切片
	for id, up := range userManager.onlineUsers{
		// 遍历用户后 id为key值
		// 如果和自身id一样时 则不通知自己
		if id == userid{
			continue
		}
		//开始通知
		up.NotifyOnline(userid)
	}
}

func (this *UserProcess) NotifyOnline(userid string)(){
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userid
	notifyUserStatusMes.Status = message.UserOnline


	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("NotifyOnline时序列化出错(小), 原因是：", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("NotifyOnline时序列化出错(大), 原因是",err)
		return
	}
	// 发送
	tf := &control.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("在线好友列表发送失败, 原因是",err)
		return
	}
}