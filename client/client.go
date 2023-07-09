package client

import (
	"LinuxProgramDesign/protocol"
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

func LoginClient(conn net.Conn) (string, error) {
	var name string
	fmt.Println("Please input your name:")
	fmt.Scanln(&name)
	if name == "Server" {
		return "", fmt.Errorf("Name can't be Server")
	}
	login := protocol.Login{Username: name}
	data, err := protocol.SerializeData(login)
	if err != nil {
		panic(err)
	}
	conn.Write(data)
	loginmsg, err := ReadOnceMessage(conn)
	if err != nil {
		panic(err)
	}

	var serverMessage protocol.Status
	protocol.UnserializeData(loginmsg, &serverMessage)
	if serverMessage.Status != 0 {
		return "", fmt.Errorf("Login failed")
	}
	return name, nil
}

func MessageLinster(conn net.Conn, username string, isDisconnect chan bool) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("[Debug] -> MessageLinster", err)
			return
		}
		if n == 0 {
			continue
		}
		fmt.Println("[Debug] -> MessageLinster", string(buf[:n]))
		error := MessageHandler(buf[:n], username, isDisconnect)
		if error != nil {
			PrintError(error)
			//fmt.Println(error)
		}
	}
}

func MessageHandler(msg []byte, username string, isDisconnect chan bool) error {
	var serverMessage protocol.SercetMessage
	var message protocol.Message
	error := protocol.UnserializeData(msg, &message)
	if error != nil {
		return error
	}
	if message.Username == "Server" {
		PrintServerMessage(message.Content)
		if message.Content == "Timeout!" {
			isDisconnect <- true
			return nil
		}
		return nil
	}
	error = protocol.UnserializeData(msg, &serverMessage)
	if error != nil {
		return error
	}
	//if serverMessage.UsernameFrom == "Server" {
	//	fmt.Println(serverMessage.Content)
	//	return nil
	//}

	if serverMessage.UsernameTo == "All" {
		PrintAllmessage(serverMessage.UsernameFrom, serverMessage.Content)
	}
	if serverMessage.UsernameTo == username {
		PrintSercetMessage(serverMessage.UsernameFrom, serverMessage.Content)
	}
	return nil
}

func SendMessageHandler(conn net.Conn, username string, isDisconnect chan bool) {
	for {

		var cmd, target string

		fmt.Scanf("%s%s", &cmd, &target)
		c := cmd
		t := target
		//fmt.Println(c, "||", t)
		//c, t, _ := parseString(cmd)

		switch c {
		case "": // wtf
			break
		case "/help":
			PrintHelp()
			break
		case "/all":
			fmt.Println("Send message to all", t)
			SendAllMessage()
			break
		case "/to":
			fmt.Println("Send message to", t)
			break
		case "/exit":
			fmt.Println("Bye!")
			isDisconnect <- true
			break
		default:
			PrintHelp()
			break
		}

	}
}

func Run() {

	isDisconnect := make(chan bool)
	//	conn := InitClient()
	//username, err := LoginClient(conn)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Login successfully, welcome", username)
	//go MessageLinster(conn, username, isDisconnect)

	// input host and port
	go SendMessageHandler(isDisconnect)
	for <-isDisconnect != true {
		break
	}

}
