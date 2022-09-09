package model

// 定义一个用户的结构体
type User struct{
	// 为了序列化 和 反序列化成功
	// 用户信息json字符串的key 和 该结构体的字段对应的tag名字一致
	UserId string `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}