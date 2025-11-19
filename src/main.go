package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(helpMsg)
		return
	}

	command := os.Args[1]

	commandHandler(command)	
}
