package main

import (
	"fmt"
	"strings"
)

func parseString(input string) (cmd, target string, err error) {
	parts := strings.SplitN(input, " ", 2)
	fmt.Println(parts)
	cmd = parts[0]
	target = parts[1]
	return
}

func main() {
	test := "aa aa"
	a, b, _ := parseString(test)

	fmt.Println(a)
	fmt.Println(b)
}
