package main

import (
	"fmt"
	"github.com/Txinkang/zinx/ziface"
	"github.com/Txinkang/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("PingRouter PreHandle")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("PreHandle ping\n"))
	if err != nil {
		fmt.Println("callback PreHandle err:", err)
	}
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("PingRouter Handle")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Handle ping\n"))
	if err != nil {
		fmt.Println("callback Handle err:", err)
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("PingRouter PostHandle")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("PostHandle ping\n"))
	if err != nil {
		fmt.Println("callback PostHandle err:", err)
	}
}

func main() {
	s := znet.NewServe("zinx 1.0.0")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
