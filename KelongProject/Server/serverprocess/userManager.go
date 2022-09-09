package serverprocess

import(
	"fmt"
)

type UserManager struct{
	onlineUsers map[string]*UserProcess
}

var(
	userManager *UserManager
)

// 对UserManager 进行初始化
func init(){
	userManager = &UserManager{
		onlineUsers : make(map[string]*UserProcess, 1024),
	}
}

//完成对 在线用户的map的 增删改查
// 增加和修改一样 对map表进行写入操作
func (this *UserManager) AddOrUpdateOnlineUser(up *UserProcess){
	this.onlineUsers[up.Process_userid] = up
}

// 删除
func (this *UserManager) DelOnlineUser(delid string){
	delete(this.onlineUsers, delid)
}

// 返回当前所有在线用户
func (this *UserManager) GetAllOnlineUsers()(map[string]*UserProcess){
	return this.onlineUsers
}

// 根据id返回对应的值
func (this *UserManager) GetOnlineUserById(userid string)(up *UserProcess, err error){
	up, ok := this.onlineUsers[userid]
	if !ok {
		err = fmt.Errorf("在server.userManager查询用户%v 不存在", userid)
		return
	}
	return
	
}