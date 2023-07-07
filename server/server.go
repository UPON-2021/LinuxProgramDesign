package server

import (
	"fmt"
	"github.com/gammazero/workerpool"
	"net"
)

// 定义全局 map 存储在线用户 key:username, value: Client
var onlineMap = make(map[string]Client)

// 定义全局 channel 处理消息
var message = make(chan []byte, 20)

// 处理客户端连接请求
func HandleConnect(conn net.Conn) {
	defer conn.Close()
	// 登录处理
	username, err := LoginHandler(conn)
	if err != nil {
		fmt.Println("[-]", "From", conn.RemoteAddr().String(), "login err:", err)
		return
	}
	message <- NEW_USER_MESSAGE(username)
	clnt := Client{make(chan []byte), username, conn.RemoteAddr().String()}
	onlineMap[username] = clnt

	go WriteMsgToClient(clnt, conn)
	go UserMsgHandler(clnt, conn)

	for {
	} //阻止线程退出
}

func Run(host string, port string, volume int) {
	// 监听端口
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	wp := workerpool.New(volume)
	go MsgManager()
	// 循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		wp.Submit(
			func() {
				HandleConnect(conn)
			})
	}
}
