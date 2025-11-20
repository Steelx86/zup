package main

import (
	"fmt"
	"os"

	"github.com/dlambda/zup/pkg/zup"
)

const (
	helpMsg  = 
`zup: zup [-OPTION]
Options:
  -n <name>     Initilize a .zup file
  -r <key>      Read zup file
  -w <text>     Writes a new entry
  -g            Generates a hash key
  -s            Host a zupfile
  -f            Fetch a zupfile from host
  -h            Displays a help message
`
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print(helpMsg)
		return
	}

	command := os.Args[1]

	switch command {
	case "-n":
		zup.NewZup(os.Args[2])
	
	case "-g":
		zup.GenerateZupKey()
	case "-s":
		return 
	case "-f":
		return
	case "-h":
		fmt.Print(helpMsg)
	default:
		fmt.Print(helpMsg)
	}
}
