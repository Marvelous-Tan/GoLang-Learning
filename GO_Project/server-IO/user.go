package main

import "net"

type User struct {
	Name    string
	Addr    string
	Conn    net.Conn// 当前用户的连接
	channel chan string
}

// 创建一个用户
func NewUser(conn net.Conn) *User {
	user := &User{
		Name:    conn.RemoteAddr().String(),// 当前用户的名称
		Addr:    conn.RemoteAddr().String(),// 当前用户的地址
		Conn:    conn,// 当前用户的连接
		channel: make(chan string),// 当前用户的channel
	}
	go user.ListenMessage()// 监听当前用户的channel，一旦有消息，就发送给对端客户端
	return user
}

// 监听当前用户的channel，一旦有消息，就发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.channel
		this.Conn.Write([]byte(msg + "\n"))// 发送消息给对端客户端
	}
}