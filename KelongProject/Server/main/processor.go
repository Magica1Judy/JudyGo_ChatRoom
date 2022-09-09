package main
import(
	"fmt"
	"net"
	"LargeProject/KelongProject/Common/control"
	"io"
	"LargeProject/KelongProject/Common/message"
	"LargeProject/KelongProject/Server/serverprocess"
)


// 创建一个结构体
type Processor struct{
	Conn net.Conn
}

func (this *Processor) process()(err error){
	//循环的客户端发送消息
	for {
		tf := &control.Transfer{
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出了，服务器端也退出")
				return err
			}else{
				fmt.Println("在Processor的ReadPkg出错了, 原因：",err)
				return err
			}
		}
		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}

}

// 根据客户端发送的消息种类来判断，调用不同功能来处理
func (this *Processor)ServerProcessMes(mes *message.Message)(err error){
	switch mes.Type {
	case message.LoginType :
		up := &serverprocess.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	
	case message.RegisterType:
		up := &serverprocess.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessRegister(mes)

	case message.SendFreeMesType:
		mesprocess := serverprocess.ServerProcess{}
		mesprocess.ServerSendGroupMes(mes)

	case message.ExitFreeRoomType:
		up := &serverprocess.UserProcess{
			Conn : this.Conn,
		}
		up.OfflineUsers(mes)

	case message.TalkToSomeBodyType:
		mesprocess := serverprocess.ServerProcess{}
		mesprocess.ServerSendSomeBodyMes(mes)
	}


	


	return


}
