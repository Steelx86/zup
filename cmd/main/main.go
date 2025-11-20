package main

import (
	"fmt"
	"os"

	"github.com/dlambda/zup/pkg/zup"
)

const (
	helpMsg  = 
`Zup 0.1 Alpha edition 2025
zup [-OPTION] <param?>
Options:
  -n <name>     Initilize a .zup file
  -r <key>      Read zup file
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

	command := os.Args[1]

	switch command {
	case "-n":
		handleNewZup(os.Args[2]) // unfinished
	case "-r":
		handleReadFile()
	case "-w":
		handleReadFile()
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

func handleNewZup(name string) {

}

func handleReadFile() {

}

func handleGenerateKey() {
	key, err := zup.GenerateZupKey()
	if err != nil {
		fmt.Printf("failed to generate key: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Generated key: %s", key)
}

func handleHost() {

}

func handlePull() {

}

