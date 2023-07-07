package main

import (
	"LinuxProgramDesign/protocol"
	"fmt"
)

func main() {
	var u protocol.Message

	u.Content = "aaaaaaaaaa"

	data, err := protocol.SerializeData(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
	var u2 protocol.Message
	err = protocol.UnserializeData(data, &u2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(u2)
	fmt.Print(u2.Content)
}
