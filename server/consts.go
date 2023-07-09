package server

import "LinuxProgramDesign/protocol"

var GreetingMessage protocol.Message = protocol.Message{
	Username: "Server",
	Content:  "Welcome to the chat room! \n Please give me your NAME.",
}

var TimeoutMessage protocol.Message = protocol.Message{
	Username: "Server",
	Content:  "Timeout!",
}

func GREETING_MESSAGE() []byte {
	data, err := protocol.SerializeData(GreetingMessage)
	if err != nil {
		return nil
	}
	return data
}

func LEAVING_MESSAGE(username string) []byte {
	LeavingMessage := protocol.Message{
		Username: "Server",
		Content:  username + "have left the chat room!",
	}
	data, err := protocol.SerializeData(LeavingMessage)
	if err != nil {
		return nil
	}
	return data
}

func TIMEOUT_MESSAGE() []byte {
	data, err := protocol.SerializeData(TimeoutMessage)
	if err != nil {
		return nil
	}
	return data
}

func NEW_USER_MESSAGE(username string) []byte {
	newUserMessage := protocol.Message{
		Username: "Server",
		Content:  username + " has joined the chat room!",
	}
	data, err := protocol.SerializeData(newUserMessage)
	if err != nil {
		return nil
	}
	return data
}

func FULL_MESSAGE() []byte {
	fullMessage := protocol.Message{
		Username: "Server",
		Content:  "The chat room is full!",
	}
	data, err := protocol.SerializeData(fullMessage)
	if err != nil {
		return nil
	}
	return data
}
