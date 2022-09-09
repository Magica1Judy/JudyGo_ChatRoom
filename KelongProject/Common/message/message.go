package message


const(
	LoginType = "Login" // 登录操作判断
	LoginReturnType = "LoginReturn"// 登录检验操作

	RegisterType = "Register" // 注册操作判断
	RegisterReturnType = "RegisterReturn" // 注册检验操作

	NotifyUserStatusMesType = "NotifyUserStatusMes"

	SendFreeMesType = "SendFreeMes"

	ExitFreeRoomType = "ExitFreeRoom"

	TalkToSomeBodyType = "TalkToSomeBody"
)
	
const(
	UserOffline = "离线"
	UserOnline = "在线"
	UserBusy = "忙碌"
)

type Message struct{
	Type string `json:"type"`
	Data string `json:"data"`
}


type LoginMes struct{
	UserId string `json:"userid"`
	UserPwd string `json:"userpwd"`
	UserName string `json:"username"`
}

type LoginReturnMes struct{
	ErrorNum int `json:"errornum"`
	ErrorData string `json:"errordata"`

	// 返回一个在线用户集合 保存已登录用户的id
	AllLoginUsers []string
}

type RegisterMes struct{
	// 为了序列化 和 反序列化成功
	// 用户信息json字符串的key 和 该结构体的字段对应的tag名字一致
	UserId string `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type RegisterReturnMes struct{
	ErrorNum int `json:"errornum"`
	ErrorData string `json:"errordata"`
}


// 为了配合服务器端 推送用户上线通知
type NotifyUserStatusMes struct{
	UserId string `json:"userid"`
	Status string `json:"status"`
}

// 群发消息
type TalkData struct{
	Content string `json:"content"`// 内容
	User // 匿名结构体
}

// 私聊消息
type SercetTalk struct{
	Content chan string
	UserId string
}

type OfflineData struct{
	UserId string `json:"userid"`
}

