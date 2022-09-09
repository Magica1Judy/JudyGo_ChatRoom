package clientprocess

import(
	"LargeProject/KelongProject/Common/message"
	"encoding/json"
	"fmt"

)

func outputMestoROOM(mes *message.Message){
	// 对返回数据进行处理 反序列化
	var receiveMes message.TalkData
	err := json.Unmarshal([]byte(mes.Data), &receiveMes)
	if err != nil {
		fmt.Println("outputMestoROOM失败, 错误原因是: ", err.Error())
		return
	}
	info := fmt.Sprintf("用户:\t%v 对大家说: \t%v",receiveMes.UserId, receiveMes.Content)
	fmt.Println(info)
	fmt.Println()
}

func outputMestoSomeBody(mes *message.Message){
	// 对返回数据进行处理 反序列化
	var receiveMes message.TalkData
	err := json.Unmarshal([]byte(mes.Data), &receiveMes)
	if err != nil {
		fmt.Println("outputMestoSomeBody失败, 错误原因是: ", err.Error())
		return
	}
	fmt.Println(receiveMes.Content)
}