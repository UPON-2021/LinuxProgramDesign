package server

import "LinuxProgramDesign/protocol"

var GreetingMessage protocol.Message = protocol.Message{
	Username: "Server",
	Content:  "Welcome to the chat room! \n Please give me your NAME.",
}

var LeavingMessage protocol.Message = protocol.Message{
	Username: "Server",
	Content:  "Leaving the chat room!",
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

func LEAVING_MESSAGE() []byte {
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
