package main

import (
	"fmt"
	"os"

	"github.com/dlambda/zup/pkg/zup"
)

const (
	helpMsg  = `Usage: zup [FILE] [KEY] 
Try 'zup help' for more information\n`
)

func main() {
	command := os.Args[1]

	switch command {
	case "-n":
		zup.newZup(os.Args[2])
	case "-h":
		fmt.Println(helpMsg)
	case "-g":
		zup.generateZupKey()
	case "-s": // sync with host
		return 
	case "-f": // forward port for hosting
		return
	default:
		fmt.Println(helpMsg)
	}
}
