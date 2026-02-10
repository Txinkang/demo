package main

import (
	"fmt"
	"github.com/Txinkang/zinx/znet"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("clientTest start")
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err: ", err)
		return
	}

	for {
		// 发封包信息
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte("zinx ClientTest Msg")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("client write err: ", err)
			return
		}

		// 取出流中的head
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("client read head err: ", err)
			break
		}
		// 拆包
		msgHead, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("client unpack err: ", err)
			return
		}
		// 看有没有数据需要取出
		if msgHead.GetDataLen() > 0 {
			// 创建message类型的变量来接收数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msgHead.GetDataLen())
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("client read msg err: ", err)
				break
			}
			fmt.Println("client read msgId	= ", msg.GetMsgId(), ", len = ", msg.GetDataLen(), ", data = ", string(msg.Data))
		}

		time.Sleep(time.Second)
	}
}
