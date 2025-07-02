package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	IP   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	// 在线用户的列表的锁
	// 什么锁？读写锁
	// 用来保护OnlineMap的读写操作，防止多个goroutine同时操作OnlineMap，导致数据不一致
	mapLock   sync.RWMutex
	// 消息广播的channel
	Message chan string
}

// 创建一个服务器
func NewServer(ip string, port int) *Server {
	return &Server{
		IP:   ip,
		Port: port,
		OnlineMap: make(map[string]*User),// 在线用户的列表
		Message: make(chan string),// 消息广播的channel
	}
}

// 广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 监听Message广播消息channel的goroutine，一旦有消息，就发送给所有在线用户
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		this.mapLock.RLock() // 读锁
		for _, user := range this.OnlineMap {
			user.channel <- msg
			fmt.Println("发送消息给", user.Name, "消息为", msg)
		}
		this.mapLock.RUnlock() // 读锁解锁
	}
}

// 处理客户端请求
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("连接建立")
	
	// 创建用户
	user := NewUser(conn)
	
	// 用户上线，将用户加入到在线用户列表中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	// 广播当前用户上线消息
	this.BroadCast(user, "已上线")
	
		// 监听用户消息
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Printf("用户 %s 断开连接\n", user.Name)
				// 用户下线，从在线用户列表中删除
				this.mapLock.Lock()
				delete(this.OnlineMap, user.Name)
				this.mapLock.Unlock()
				// 广播用户下线消息
				this.BroadCast(user, "已下线")
				return
			}
			if err != nil {
				fmt.Printf("用户 %s 连接错误: %v\n", user.Name, err)
				// 用户下线，从在线用户列表中删除
				this.mapLock.Lock()
				delete(this.OnlineMap, user.Name)
				this.mapLock.Unlock()
				// 广播用户下线消息
				this.BroadCast(user, "已下线")
				return
			}
			
			// 提取用户消息（去除\n）
			msg := string(buf[:n-1])
			
			// 广播消息
			this.BroadCast(user, msg)
		}
	}()
	
	// 阻塞当前goroutine
	select {}
}

// 启动服务器
func (this *Server) Start() {
	// 1. 监听端口（socket listen）
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// 启动监听Message广播消息的goroutine
	go this.ListenMessage()

	// 2. 接收客户端请求（accept）
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}
		// 3. 处理客户端请求（do handler）
		go this.Handler(conn)
	}

	// 4. 关闭连接 (close listen socket)
	defer listener.Close()
}

