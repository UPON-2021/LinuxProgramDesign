package main

import (
	"LinuxProgramDesign/server"
	"fmt"
)

func main() {
label:
	var volume int
	fmt.Println("Please input the volume of worker pool:")
	_, err := fmt.Scanln(&volume)
	if err != nil {
		fmt.Println("Please input a number")
		goto label
	}
	server.Run("0.0.0.0", "8848", volume)
}
