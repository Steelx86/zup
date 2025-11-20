package zup

import (
	"fmt"
	"os"
	"strings"
	"encoding/hex"

	"github.com/dlambda/zup/internal/encryption"
)

const (
	KEY_SIZE = 32
)

func newZup(name string) {
	if strings.HasSuffix(name, ".zup") {
		os.Create(name)
	} else {
		os.Create(name + ".zup")
	}

	key, err := encryption.generateKey(KEY_SIZE)
	if err != nil {
		fmt.Printf("Key generation error: %v\n", err)
	}

	readableKey := hex.EncodeToString(key)

	fmt.Printf("The key for your new zup file is: %s\n", readableKey)
}

func openZup(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
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
