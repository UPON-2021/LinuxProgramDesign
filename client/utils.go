package client

import (
	"LinuxProgramDesign/protocol"
	"fmt"
	"net"
	"strings"
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

// format /[cmd] [target] [message]
func parseString(input string) (cmd, message string, err error) {
	fmt.Println("input:", input)
	parts := strings.SplitN(input, " ", 2)
	fmt.Println(parts)
	cmd = parts[0]
	message = parts[1]
	return
}

func SendAllMessage(conn net.Conn, username, message string) error {
	var msg protocol.SercetMessage
	msg.UsernameFrom = username
	msg.UsernameTo = "All"
	msg.Content = message
	data, err := protocol.SerializeData(msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func SendSercetMessage(conn net.Conn, username, to, message string) error {
	var msg protocol.SercetMessage
	msg.UsernameFrom = username
	msg.UsernameTo = to
	msg.Content = message
	data, err := protocol.SerializeData(msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func SendLeaveMessage(conn net.Conn, username string) error {
	var msg protocol.SercetMessage
	msg.UsernameFrom = "Server"
	msg.UsernameTo = "All"
	msg.Content = username + " leave the chatroom"
	data, err := protocol.SerializeData(msg)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}
