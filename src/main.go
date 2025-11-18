package main

import (
	"fmt"
	"os"
)

const (
	KEY_SIZE = 32
)

func main() {
	command := os.Args[1]

	switch command {
	case "open":
		openZup(os.Args[2])
	case "new":
		newZup(os.Args[2])
	case "help":
		fmt.Println(helpMsg)
	case "generate":
		generateZupKey()
	default:
		fmt.Println(helpMsg)
	}
}
