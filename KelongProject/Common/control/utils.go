package control

import(
	"net"
	"LargeProject/KelongProject/Common/message"
	"fmt"
	"encoding/binary"
	"encoding/json"
)

type Transfer struct{
	Conn net.Conn
	Buf [8096]byte
}

func (this *Transfer) ReadPkg()(mes message.Message, err error){
	// fmt.Println("线程开始读取信息...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("Utils的ReadPkg中的conn.Read(1) 读取失败, 原因是：", err)
		return
	}
	//fmt.Println(" ReadPkg() 读到的信息buf: ", this.Buf[:4])

	// 
	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(this.Buf[:4])

	n, err := this.Conn.Read(this.Buf[:pkglen])
	if n != int(pkglen) || err != nil {
		fmt.Println("读取返回的验证内容时出现错误, 原因是：", err)
		return
	}

	err = json.Unmarshal(this.Buf[:pkglen], &mes)
	if err != nil{
		fmt.Println("在ReadPkg()中反序列化失败, 原因是：", err)
		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte)(err error){
	// 要把一个长度发出去 现在只得到了一个切片
	// 首先用LEN获取到切片的长度 后利用方法将它进行转换成byte型
	var datalen uint32 = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], datalen)

	// 获得了一个byte切片 即可进行传入操作
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("WritePkg传输长度有误, 原因：",err)
		return
	} 

	n, err = this.Conn.Write(data)
	if n != int(datalen) || err != nil {
		fmt.Println("WritePkg传输内容有误, 原因：",err)
		return
	} 
	return


}