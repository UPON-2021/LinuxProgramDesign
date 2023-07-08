package server

import (
	"LinuxProgramDesign/protocol"
	"fmt"
	"github.com/gammazero/workerpool"
	"net"
	"time"
)

// 定义全局 channel 处理消息
var message = make(chan []byte, 20)

// 处理客户端连接请求
func HandleConnect(conn net.Conn, onlineMap2 onlineMap, isAccept chan bool) {
	defer conn.Close()
	isAccept <- true
	// 登录处理
	username, err := LoginHandler(conn)
	if err != nil {
		fmt.Println("[-]", "From", conn.RemoteAddr().String(), "login err:", err)
		return
	}
	message <- NEW_USER_MESSAGE(username)
	clnt := Client{make(chan []byte), username, conn.RemoteAddr().String()}
	AddMap(&onlineMap2, username, clnt)
	hasData := make(chan bool)
	isConnectionLost := make(chan bool)

	go WriteMsgToClient(clnt, conn)
	go UserMsgHandler(clnt, conn, hasData, isConnectionLost)

	for {
		select {
		case <-hasData:
			break
		case <-isConnectionLost: // 掉线处理，没写鉴权，没法做重连
			msg, _ := protocol.MakeMsg("[Server]", clnt.Name+" leave")
			DelMap(&onlineMap2, clnt.Name)
			message <- msg
			return
		case <-time.After(60 * time.Second):
			msg, _ := protocol.MakeMsg("[Server]", clnt.Name+" time out leave")
			DelMap(&onlineMap2, clnt.Name)
			message <- msg
			conn.Write(TIMEOUT_MESSAGE()) // 通知当前用户断开连接
			return                        // 结束当前应用
		}

	} //阻止线程退出
}

func Run(host string, port string, volume int) {
	// 监听端口
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	var onlinemap = onlineMap{
		Clients: make(map[string]Client),
	}
	wp := workerpool.New(volume)
	go MsgManager(onlinemap)
	// 循环监听客户端连接请求

	for {

		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		isAccept := make(chan bool)

		timeout := 3 * time.Second
		go func() {
			wp.Submit(
				func() {
					HandleConnect(conn, onlinemap, isAccept)
				})
		}()

		// 使用select语句同时监听连接和超时信号
		select {
		case <-isAccept:
			for {
			}
		case <-time.After(timeout):
			// 超时，关闭连接
			conn.Close()
			return
		}

	}
}
