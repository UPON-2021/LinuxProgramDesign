package client

import (
	"LinuxProgramDesign/protocol"
	"fmt"
)

func PrintLoginMessage(message []byte) {
	var serverMessage protocol.Message
	protocol.UnserializeData(message, &serverMessage)
	fmt.Println("[*]", serverMessage.Username, ":", serverMessage.Content)
}

func PrintAllmessage(username string, message string) {
	fmt.Println("[* Public ]", username, ":", message)
}

func PrintSercetMessage(username string, message string) {
	fmt.Println("[* Secret ]", username, ":", message)
}

func PrintServerMessage(message string) {
	fmt.Println("[* Server ]", message)
}

func PrintError(err error) {
	fmt.Println("[* Error ]", err)
}
