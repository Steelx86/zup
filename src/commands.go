package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func commandHandler(command string) {
	switch command {
	case "open":
		openZup(os.Args[2])
	case "new":
		newZup(os.Args[2])
	case "help":
		fmt.Println(helpMsg)
	case "generate":
		generateZupKey()
	case "sync": // sync with host
		return // place holder
	case "host": // open port for hosting
		return
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println(helpMsg)
	}
}

func REPL() {
	fmt.Print("Welcome to Zup alpha!\n Type '.help' for more information\n")

	var input string

	for input != ".exit" {
		fmt.Print("zup> ")
		fmt.Scanln(&input)
	}
}

func openZup(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
}

func newZup(name string) {
	if strings.HasSuffix(name, ".zup") {
		os.Create(name)
	} else {
		os.Create(name + ".zup")
	}

	key, err := generateKey(KEY_SIZE)
	if err != nil {
		fmt.Printf("Key generation error: %v\n", err)
	}

	readableKey := hex.EncodeToString(key)

	fmt.Printf("The key for your new zup file is: %s\n", readableKey)
}

func generateZupKey() {
	key, err := generateKey(KEY_SIZE)
	if err != nil {
		fmt.Printf("Key generation error: %v\n", err)
		return
	}

	readableKey := hex.EncodeToString(key)

	fmt.Printf("Generated key: %s\n", readableKey)
}


