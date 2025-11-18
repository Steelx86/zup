package src

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]

	switch command {
		case "open": openZup()
		case "new": newZup()
		case "help": fmt.Println(helpMsg)
		default: fmt.Println(helpMsg)
	}

}
