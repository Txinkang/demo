package main

import (
	"fmt"
	"github.com/Txinkang/zinx/znet"
	"io"
	"net"
)

func main() {
	// 开启服务创建监听
	listenner, err := net.Listen("tcp", ":7777")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	// 开始阻塞接受连接
	for {
		// 接受连接
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}

		// 处理客户端请求
		go func(conn net.Conn) {
			// 创建封包拆包对象
			dp := znet.NewDataPack()

			// 申请空间获取数据
			for {
				// 1、先读出流中的head部分
				headData := make([]byte, dp.GetHeadLen())
				_, err = io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("Error reading:", err.Error())
					break
				}

				// 2、拆包到msg
				msgHead, err := dp.UnPack(headData)
				if err != nil {
					fmt.Println("Error unpacking:", err.Error())
					return
				}

				// 3、读取字节流数据
				if msgHead.GetDataLen() > 0 {
					msg := msgHead.(*znet.Message)
					msg.Data = make([]byte, msg.GetDataLen())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("Error UnPack reading:", err.Error())
						return
					}
					fmt.Println("Msg id =:", msg.Id, ", len = ", msg.DataLen, ", data = ", string(msg.Data))
				}
			}
		}(conn)
	}
}
