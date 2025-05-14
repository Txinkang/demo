package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	Chan   chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		Chan:   make(chan string),
		conn:   conn,
		server: server,
	}

	go user.ListenMsg()
	return user
}

// 用户上线
func (user *User) UserOnline() {
	//把新连接的user加入到map中
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	//设置广播消息
	user.server.Boradcast(user, "已上线")
}

// 用户下线
func (user *User) UserOffline() {
	//把user从map中删除
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	//设置广播消息
	user.server.Boradcast(user, "已下线")
}

// 处理消息
func (user *User) MsgHandler(msg string) {
	switch {
	case msg == "who":
		user.server.mapLock.Lock()
		for _, queryUser := range user.server.OnlineMap {
			sendMsg := "[" + queryUser.Addr + "] :" + queryUser.Name + ":在线\n"
			user.SendMsg(sendMsg)
		}
		user.server.mapLock.Unlock()
	case strings.HasPrefix(msg, "rename|"):
		newName := strings.Split(msg, "|")[1]
		if _, ok := user.server.OnlineMap[newName]; ok {
			user.SendMsg("该用户名已被注册")
		} else {
			//先删除旧的，再添加新的
			user.server.mapLock.Lock()
			delete(user.server.OnlineMap, user.Name)
			user.server.OnlineMap[newName] = user
			user.server.mapLock.Unlock()

			user.Name = newName
			user.SendMsg("当前用户名已更新\n")
		}
	case strings.HasPrefix(msg, "to|"):
		//获取用户名
		queryUser := strings.Split(msg, "|")[1]
		//检查用户是否存在
		if targetUser, ok := user.server.OnlineMap[queryUser]; ok {
			content := strings.Split(msg, "|")[2]
			targetUser.SendMsg(user.Name + "对您说：" + content)
		}
	default:
		user.server.Boradcast(user, msg)
	}
}

// 给具体用户发消息
func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}
func (user *User) ListenMsg() {
	for {
		msg := <-user.Chan
		user.conn.Write([]byte(msg + "\n"))
	}
}
