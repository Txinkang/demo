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

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Handle ping\n"))
	if err != nil {
		fmt.Println("callback Handle err:", err)
	}
}

func main() {
	s := znet.NewServe("zinx 1.0.0")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
