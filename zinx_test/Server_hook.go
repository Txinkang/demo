package main

import (
	"fmt"
	"github.com/Txinkang/zinx/ziface"
	"github.com/Txinkang/zinx/znet"
)

type PingRouterHook struct {
	znet.BaseRouter
}

// Test Handle
func (this *PingRouterHook) Handle(request ziface.IRequest) {
	fmt.Println("PingRouter Handle")
	// 读取客户端数据
	fmt.Println("recv from client: msgId = ", request.GetMsgId(), "data = ", string(request.GetData()))

	// 回写数据
	err := request.GetConnection().SendMsg(0, []byte("PingRouter ping...ping..."))
	if err != nil {
		fmt.Println("Error sending data: ", err.Error())
	}
}

type HelloZinxRouterHook struct {
	znet.BaseRouter
}

// Test Handle
func (this *HelloZinxRouterHook) Handle(request ziface.IRequest) {
	fmt.Println("HelloZinxRouter Handle")
	// 读取客户端数据
	fmt.Println("recv from client: msgId = ", request.GetMsgId(), "data = ", string(request.GetData()))

	// 回写数据
	err := request.GetConnection().SendMsg(1, []byte("HelloZinxRouter ping...ping..."))
	if err != nil {
		fmt.Println("Error sending data: ", err.Error())
	}
}

// 创建连接时执行
func DoConnStart(conn ziface.IConnection) {
	fmt.Println("DoConnStart ok")
	conn.SendMsg(0, []byte("DoConnStart ok"))
}

// 断开连接时执行
func DoConnStop(conn ziface.IConnection) {
	fmt.Println("DoConnStart ok")
	conn.SendMsg(0, []byte("DoConnStart ok"))
}
func main() {
	s := znet.NewServe()
	s.SetOnConnStart(DoConnStart)
	s.SetOnConnStop(DoConnStop)
	s.AddRouter(0, &PingRouterHook{})
	s.AddRouter(1, &HelloZinxRouterHook{})
	s.Serve()
}
