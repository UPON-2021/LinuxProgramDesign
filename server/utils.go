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
	for {
		for msg := range clnt.Messages {
			fmt.Println("WriteMsgToClient:", clnt.Addr, clnt.Name, string(msg))
			conn.Write(msg)
		}
	}
}

func MsgManager() {
	for {
		fmt.Println("MsgManager: waiting for data")
		msg := <-message
		fmt.Println("MsgManager", msg)
		for _, clnt := range onlineMap {
			clnt.Messages <- msg
		}
	}
}

func UserMsgHandler(clnt Client, conn net.Conn, hasData chan bool) {
	// 检测用户是否有消息发送
	buf := make([]byte, 2048)
	for {
		fmt.Println("UserMsgHandler: waiting for data")
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println("UserMsgHandler:", buf[:n])
		message <- buf[:n]
		hasData <- true
	}
}
