package main

import (
	"fmt"
	"github.com/Txinkang/zinx/ziface"
	"github.com/Txinkang/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("PingRouter Handle")
	// 读取客户端数据
	fmt.Println("recv from client: msgId = ", request.GetMsgId(), "data = ", string(request.GetData()))

	// 回写数据
	err := request.GetConnection().SendMsg(1, []byte("Handle ping...ping..."))
	if err != nil {
		fmt.Println("Error sending data: ", err.Error())
	}
}

func main() {
	s := znet.NewServe("zinx 1.0.0")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
