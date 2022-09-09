package main

import(
	"fmt"
	"net"
	"time"
	"LargeProject/KelongProject/Server/model"

)

// 做初始化
func init(){
	//服务器启动时，我们就进行初始化
	initPool("127.0.0.1:6379", 8, 0, 300 * time.Second)
	model.MyUserDao = model.NewUserDao(pool)
}

// server创建的链接线程
func process(conn net.Conn){
	defer conn.Close()
	//创建一个总控
	processor := &Processor{
		Conn : conn,
	}
	err := processor.process()
	if err != nil {
		fmt.Println("**main中协程 客户端和服务端通讯协程**有问题: ",err)
		return
	}
	
}




func main(){
	fmt.Println("************聊天服务器开始监听了**************")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("在Server的main中, 与服务器的链接 出现问题: ", err)
		return
	}

	// 建立与客户端的持续链接
	for {
		fmt.Println("聊天服务器等待服务端的链接中。。。")
		conn, err := listen.Accept()
		if err != nil {
			// 不做return 因为可能有其它客户端连接
			fmt.Println("在Server的main中, 服务器提供的 连接接口 出现问题: ", err)
		}else{
			fmt.Printf("\n用户地址为[ %v ]与服务端的连接成功了\n", conn.RemoteAddr().String())
		}
		go process(conn)

	}

}