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
  -n <name>        Initilize a .zup file
  -r <name> <key>  Read zup file
  -d               Declare data format
  -w <text>        Writes a new entry
  -g               Generates a hash key
  -s               Host a zupfile
  -p <pull>        Pull the zupfile from host
  -h               Displays a help message
`
)

func main() {
	commands := os.Args

	if len(os.Args) < 2 {
		fmt.Print(helpMsg)
		return
	}


	switch commands[1] {
	case "-n":
		if len(commands) < 3 {
			fmt.Println("Please provide a name for the zup file.\nEX: zup zupfile.zup")
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
		if len(commands) < 4 {
			fmt.Println("Please provide a name and a key\nEX: zup zupfile.zup <KEY>")
			return
		}

		name := commands[2]
		key := commands[3]

		zupData, err := zup.OpenZup(name, key)
		if err != nil {
			fmt.Printf("Zupfile failed to open: %v", err)
		}

		fmt.Println(zupData.String())
	case "-d":
		
	case "-w":
		// unfinished
	case "-g":
		key, err := zup.GenerateZupKey()
		if err != nil {
			fmt.Printf("failed to generate key: %v", err)
			return
		}

		fmt.Println("Generated key:", key)
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

