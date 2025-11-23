package main

import (
	"fmt"
	"os"

	"github.com/dlambda/zup/pkg/zup"
)

const (
	helpMsg = `Zup 0.1 Alpha edition 2025
zup [-OPTION] <param?>
Options:
  -n <name>     Initilize a .zup file
  -r <key>      Read zup file
  -d            Declare data format
  -w <text>     Writes a new entry
  -g            Generates a hash key
  -s            Host a zupfile
  -p <pull>     Pull the zupfile from host
  -h            Displays a help message
`
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print(helpMsg)
		return
	}

	commands := os.Args

	switch commands[1] {
	case "-n":
		if len(os.Args) < 3 {
			fmt.Print("Please provide a name for the zup file.\n")
			return
		}

		name := commands[2]

		key, err := zup.InitZup(name)
		if err != nil {
			fmt.Printf("failed to initialize zup file: %v", err)
			return
		}

		fmt.Printf("Initialized zup file. Key: %s\n", key)
	case "-r":
		// unfinished
	case "-d":
		// unfinished
	case "-w":
		// unfinished
	case "-g":
		key, err := zup.GenerateZupKey()
		if err != nil {
			fmt.Printf("failed to generate key: %v", err)
			return
		}

		fmt.Printf("Generated key: %s", key)
	case "-s":
		// unfinished
	case "-p":
		// unfinished
	case "-h":
		fmt.Print(helpMsg)
	default:
		fmt.Print(helpMsg)
	}
}

