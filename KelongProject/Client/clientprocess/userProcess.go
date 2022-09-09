package clientprocess
import(
	"net"
	"fmt"
	"LargeProject/KelongProject/Common/message"
	"encoding/json"
	"LargeProject/KelongProject/Common/control"
	"os"
	"time"
)


type UserProcess struct{

}



func (this *UserProcess) Login(loginid string, loginpwd string)(err error){
	// 1、连接到服务器
	// 首先 进行 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil{
		fmt.Println("UserProcess的Login对服务器端拨号有误了:",err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2、——————————————————————准备发送的消息————————————————————
	var mes message.Message
	mes.Type = message.LoginType

	//创建结构体
	var loginmes message.LoginMes
	loginmes.UserId = loginid
	loginmes.UserPwd = loginpwd

	// 序列化小的结构体
	data, err := json.Marshal(loginmes)
	if err != nil {
		fmt.Println("在client.userprocess中的Login序列化出问题了, 问题是：",err)
		return
	}
	mes.Data = string(data)

	// 序列化大的结构体
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("在client.userprocess中的Login序列化出问题了, 问题是：",err)
		return
	}
	//得到了一个要发送的完整数据  一个byte切片
	//2、——————————————————————准备发送的消息————————————————————

	// 防止丢包 需要先发送一个长度进行校验
	tf := &control.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("在userProcess的Login方法发送登录内容失败, 失败原因：",err)
		return
	} 
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("在userProcess的Login方法读取返回检验结果失败, 失败原因：",err)
		return
	} 

	var loginReturnMes message.LoginReturnMes
	err = json.Unmarshal([]byte(mes.Data), &loginReturnMes)
	if loginReturnMes.ErrorNum == 200 {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>登录成功")
		
		LoginedUser.Conn = conn
		LoginedUser.UserId = loginid
		LoginedUser.UserStatus = message.UserOnline


		for _, value := range loginReturnMes.AllLoginUsers{
			if value == loginid{
				continue
			}
			// fmt.Println("用户id : ", value)

			// 对userManager.ShowClientOnlineUsers 进行初始化
			user := &message.User{
				UserId : value,
			}
			ShowClientOnlineUsers[value] = user

		}
		
		time.Sleep(time.Second)
		// 一直和服务端保持联系
		go serverProcessMes(conn)
		time.Sleep(time.Second)
		ShowMenu()
		
	}else{
		fmt.Println("登录失败")
	}
	return

}

func (this *UserProcess) Register(registerid , registerpwd ,registername string )(err error){
	// 1、连接到服务器
	// 首先 进行 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil{
		fmt.Println("clientProcess的Register注册时，链接服务器端口有误了")
		return
	}
	//延时关闭
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegisterType

	// 3、创建一个RegisterMes的结构体 
	var Registermes message.RegisterMes
	Registermes.UserId = registerid
	Registermes.UserPwd = registerpwd
	Registermes.UserName = registername

	data, err := json.Marshal(Registermes)
	if err != nil {
		fmt.Println("clientProcess的Registermes序列化出问题了, 问题是：",err)
		return
	}
	mes.Data = string(data)

	// 6、将 mes进行序列化			(大)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("clientProcess的大mes序列化出问题了, 问题是：",err)
		return
	}
	// 创建一个Transfer实例
	tf := &control.Transfer{
		Conn : conn,
	}
	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("clientProcess的注册发送的信息错误, 原因是：", err)
	}


	// 进行读取 检验是否注册成功
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("注册的读取内容失败, 失败原因：",err)
		return
	} 
	// 这是读取返回值 看是否注册成功
	// 将mes的Data部分反序列化成RegisterResMes
	var registerResMes message.RegisterReturnMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.ErrorNum == 200 {

		fmt.Println("注册成功, 您可继续重新登录")
		os.Exit(0)

		
	}else{
		fmt.Println("注册失败, 失败原因：",registerResMes.ErrorData)
		os.Exit(0)
	} 
	return
}

