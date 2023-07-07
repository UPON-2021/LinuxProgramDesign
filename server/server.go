package server

import (
	"net"
)

type Client struct {
	Messages chan string // 发送消息的通道
	Name     string      // 用户名
	Addr     string      // 网络地址：ip+port
}

// 定义全局 map 存储在线用户 key:username, value: Client
var onlineMap = make(map[string]Client)

// 定义全局 channel 处理消息
var message = make(chan string, 20)

// 处理客户端连接请求
func HandleConnect(conn net.Conn) {
	defer conn.Close()
	//clnt := Client{make(chan string), "NewUser", conn.RemoteAddr().String()}
	conn.Write(GREETING_MESSAGE())
}

func Run(host string, port string) {
	// 监听端口
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// 循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 处理客户端连接请求
		go HandleConnect(conn)
	}
}
