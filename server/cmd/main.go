package main

import "LinuxProgramDesign/server"

func main() {
	server.Run("0.0.0.0", "8848", 5)
}
