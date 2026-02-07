package main

import (
	"fmt"
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
		_, err := conn.Write([]byte("hello zinx"))
		if err != nil {
			fmt.Println("client write err: ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client read err: ", err)
			return
		}
		fmt.Printf("server call back: %s, cnt = %d\n", buf[:cnt], cnt)

		time.Sleep(time.Second)
	}
}
