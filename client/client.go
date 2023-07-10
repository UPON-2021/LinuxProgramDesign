package client

import (
	"LinuxProgramDesign/protocol"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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
			//fmt.Println("[Debug] -> MessageLinster", err)
			return
		}
		if n == 0 {
			continue
		}
		//fmt.Println("[Debug] -> MessageLinster", string(buf[:n]))
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

		reader := bufio.NewReader(os.Stdin)
		result, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("获取失败")
		}

		command := string(result) + " " // 为了防止用户输入的命令不带参数，导致的数组越界
		if command == "/exit" {
			isDisconnect <- true
			return
		}

		//var cmd, target, message string
		_tmp := strings.SplitN(command, " ", 2)
		c := _tmp[0]
		t := _tmp[1]
		//fmt.Scanf("%s%s%s", &cmd, &target, &message)

		//fmt.Println(c, "||", t)
		//c, t, _ := parseString(cmd)

		switch c {
		case "": // 奇异的bug，不知如何产生的，就这么简单修修吧(
			break
		case "/help":
			PrintHelp()
			break
		case "/all":
			SendAllMessage(conn, username, t)
			break
		case "/to":
			a := strings.SplitN(t, " ", 2)
			SendSercetMessage(conn, username, a[0], a[1])
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
	conn := InitClient()
	username, err := LoginClient(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Login successfully, welcome", username)
	go MessageLinster(conn, username, isDisconnect)

	// input host and port
	go SendMessageHandler(conn, username, isDisconnect)
	for <-isDisconnect != true {
		break
	}

}
