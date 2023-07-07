package server

import (
	"LinuxProgramDesign/protocol"
	"fmt"
	"net"
)

type Client struct {
	Messages chan []byte // 发送消息的通道
	Name     string      // 用户名
	Addr     string      // 网络地址：ip+port
}

func LoginHandler(conn net.Conn) (string, error) {
	conn.Write(GREETING_MESSAGE())
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}
	var login protocol.Login
	err = protocol.UnserializeData(buf[:n], &login)
	if err != nil {
		return "", err
	}
	fmt.Println("[+]", login.Username+" login")
	return login.Username, nil
}

func WriteMsgToClient(clnt Client, conn net.Conn) {
	for msg := range clnt.Messages {
		conn.Write(msg)
	}
}

func MsgManager() {
	for {
		msg := <-message
		for _, clnt := range onlineMap {
			clnt.Messages <- msg
		}
	}
}

func UserMsgHandler(clnt Client, conn net.Conn) {
	hasData := make(chan bool) // 检测用户是否有消息发送
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		message <- buf[:n]
		hasData <- true
	}
	//for {
	//	select {
	//	case <-hasData:
	//	case <-time.After(60 * time.Second):
	//		msg, _ := protocol.MakeMsg("[Server]", clnt.Name+" time out leave")
	//		message <- msg
	//		conn.Write(TIMEOUT_MESSAGE()) // 通知当前用户断开连接
	//		return                        // 结束当前应用
	//	}
	//}
}
