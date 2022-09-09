package model

import(
	"github.com/gomodule/redigo/redis"
	"encoding/json"
	"fmt"
	"LargeProject/KelongProject/Common/message"

)

type UserDao struct{
	pool *redis.Pool
}

var(
	MyUserDao *UserDao
)

// 工厂模式
func NewUserDao(pool *redis.Pool)(userDao *UserDao){
	userDao = &UserDao{
		pool : pool,
	}
	return
}

func (this *UserDao)GetUserById(conn redis.Conn, id string)(user *User, err error){
	res, err := redis.String(conn.Do("Hget","users",id))
	if err != nil{
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("在userDao中,将redis中的数据反序列化出错, 原因是：",err)
		return
	}
	return
}


// 登录校验 用户验证 若id和pwd都正确 则返回一个user实例
func (this *UserDao)Login(userid string, userpwd string)(user *User, err error){
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.GetUserById(conn, userid)
	if err != nil{
		return
	}
	if user.UserPwd != userpwd{
		err = ERROR_USER_PWD
		return
	}
	return
}



func (this *UserDao) Register(user *message.RegisterMes)( err error){
	//先从UserDao 的连接池中 取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	judge, _ := this.GetUserById(conn, user.UserId)
	if judge != nil {
		err = ERROR_USER_EXISTS
		return
	}
	// 这时， 说明id在redis还没有，则可以完成注册
	data, err := json.Marshal(user)
	if err != nil{
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil{
		fmt.Println("保存注册用户错误, 错误原因：", err)
		return
	}
	
	return
}

