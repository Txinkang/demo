package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	Msg       chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Msg:       make(chan string),
	}
	return server
}

// 广播消息
func (this *Server) Boradcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Msg <- sendMsg
}

// 监听Msg管道，有消息了就发送给所有用户
func (this *Server) ListenMsg() {
	for {
		msg := <-this.Msg
		this.mapLock.Lock()
		for _, user := range this.OnlineMap {
			user.Chan <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Handler(conn net.Conn) {
	//创建user
	user := NewUser(conn, this)
	//上线
	user.UserOnline()
	isAlive := make(chan bool)

	//接收用户消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				//下线
				user.UserOffline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("无法读取客户端连接消息：", err)
				return
			}

			//处理消息
			msg := string(buf[:n-1])
			user.MsgHandler(msg)
			isAlive <- true
		}
	}()
	//阻塞Handler，让它一直运行
	for {
		select {
		case <-isAlive:

		case <-time.After(time.Second * 300):
			user.SendMsg("您已被踢出群聊")
			close(user.Chan)
			conn.Close()
			return
		}

	}
}

func (this *Server) Start() {
	// 监听服务
	listener, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println(err)
		return
	}

	//关闭监听
	defer listener.Close()

	go this.ListenMsg()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener Accept err:", err)
		}

		go this.Handler(conn)
	}
}
