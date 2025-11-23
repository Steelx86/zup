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
	commands := os.Args

	switch commands[1] {
	case "-n":
		handleNewZup(commands)
	case "-r":
		handleReadFile(commands)
	case "-w":
		handleReadFile(commands)
	case "-g":
		handleGenerateKey()
	case "-s":
		handleHost() // unfinished
	case "-p":
		handlePull() // unfinished
	case "-h":
		fmt.Print(helpMsg)
	default:
		fmt.Print(helpMsg)
	}
}

func handleNewZup(name []string) {
	if len(os.Args) < 3 {
		fmt.Print("Please provide a name for the zup file.\n")
		return
	}

	key, err := zup.InitZup(name[2])
	if err != nil {
		fmt.Printf("failed to initialize zup file: %v", err)
		return
	}

	fmt.Printf("Initialized zup file. Key: %s\n", key)
}

func handleReadFile(name []string) {
	if len(os.Args) < 3 {
		fmt.Print("Please provide text to write to the zup file.\n")
		return
	}
}

func handleGenerateKey() {
	key, err := zup.GenerateZupKey()
	if err != nil {
		fmt.Printf("failed to generate key: %v", err)
		return
	}

	fmt.Printf("Generated key: %s", key)
}

func handleHost() {

}

func handlePull() {

}
