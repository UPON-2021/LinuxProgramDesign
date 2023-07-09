package client

import (
	"fmt"
	"net"
)

func ReadOnceMessage(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 2048)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("ReadMessage error:", err)
		panic(err)
	}
	return buf[:n], nil
}

func ReadMessage(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("ReadMessage error:", err)
			panic(err)
		}
		return buf[:n], nil
	}
}
