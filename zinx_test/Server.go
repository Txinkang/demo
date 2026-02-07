package main

import "github.com/Txinkang/zinx/znet"

func main() {
	s := znet.NewServe("zinx 1.0.0")
	s.Serve()
}
