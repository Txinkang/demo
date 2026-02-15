package main

import (
	"fmt"
	"github.com/Txinkang/zinx/ziface"
	"github.com/Txinkang/zinx/znet"
)

type PingRouterProperty struct {
	znet.BaseRouter
}

// Test Handle
func (this *PingRouterProperty) Handle(request ziface.IRequest) {
	fmt.Println("PingRouter Handle")
	// 读取客户端数据
	fmt.Println("recv from client: msgId = ", request.GetMsgId(), "data = ", string(request.GetData()))

	// 回写数据
	err := request.GetConnection().SendMsg(0, []byte("PingRouter ping...ping..."))
	if err != nil {
		fmt.Println("Error sending data: ", err.Error())
	}
}

type HelloZinxRouterProperty struct {
	znet.BaseRouter
}

// Test Handle
func (this *HelloZinxRouterProperty) Handle(request ziface.IRequest) {
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
func DoConnStart1(conn ziface.IConnection) {
	fmt.Println("DoConnStart ok")
	err := conn.SendMsg(0, []byte("DoConnStart ok"))
	if err != nil {
		fmt.Println("Error sending data: ", err.Error())
	}
	conn.SetProperty("name", "Tang")
}

// 断开连接时执行
func DoConnStop1(conn ziface.IConnection) {
	fmt.Println("DoConnStop ok")
	if value, err := conn.GetProperty("name"); err == nil {
		fmt.Println("conn property name:", value)
	}
}
func main() {
	s := znet.NewServe()
	s.SetOnConnStart(DoConnStart1)
	s.SetOnConnStop(DoConnStop1)
	s.AddRouter(0, &PingRouterProperty{})
	s.AddRouter(1, &HelloZinxRouterProperty{})
	s.Serve()
}
