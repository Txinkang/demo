package main

import (
	"fmt"
	"github.com/Txinkang/zinx/znet"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:7777")
	if err != nil {
		fmt.Println("Error Dial:", err.Error())
		return
	}

	//创建封包拆包对象
	dp := znet.NewDataPack()

	// 创建msg1
	msg1 := &znet.Message{
		Id:      0,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("Error Pack:", err.Error())
		return
	}
	// 创建msg2
	msg2 := &znet.Message{
		Id:      0,
		DataLen: 6,
		Data:    []byte{'w', 'o', 'r', 'l', 'd', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("Error Pack:", err.Error())
		return
	}

	// 连接消息
	sendData1 = append(sendData1, sendData2...)

	// 发送消息
	conn.Write(sendData1)

	// 客户端阻塞
	//select {}
	for {
	}
}
