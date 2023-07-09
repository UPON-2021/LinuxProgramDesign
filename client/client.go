package client

import (
	"fmt"
	"net"
)

type TcpClient struct {
	name    string
	conn    net.Conn
	message chan []byte
}

func InitClient() net.Conn {
	var host string
	var port string

	fmt.Println("Please input host:")
	fmt.Scanln(&host)

	fmt.Println("Please input port:")
	fmt.Scanln(&port)

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		panic(err)
	}
	return conn

}

func LoginClient(conn net.Conn) {
	var name string
	fmt.Println("Please input your name:")
	fmt.Scanln(&name)
	conn.Write([]byte(name))
}

func Run() {
	InitClient()

	// input host and port

}
