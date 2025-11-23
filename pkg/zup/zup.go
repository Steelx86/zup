package zup

import (
	"encoding/json"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const (
	KEY_SIZE = 32
)

func InitZup(name string) (string, error) {
	if strings.HasSuffix(name, ".zup") {
		os.Create(name)
	} else {
		os.Create(name + ".zup")
	}

	key, err := GenerateKey(KEY_SIZE)
	if err != nil {
		return "", err
	}

	readableKey := hex.EncodeToString(key)

	return readableKey, nil
}

func OpenZup(fileName string, key string) (Zup, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return Zup{}, err
	}

	encryptedContent := string(file)
	content, err := decrypt(encryptedContent, []byte(key))
	if err != nil {
		return Zup{}, err
	}

	zupData, err := ParseZupString(content)
	if err != nil {
		return Zup{}, err
	}

	return zupData, nil
}

func ParseZupString(zupString string) (*Zup, error) {
	var zup Zup
	err := json.Unmarshal([]byte(zupString), &zup)
	if err != nil {
		return &Zup{}, fmt.Errorf("failed to parse Zup JSON: %v", err)
	}
	return &zup, nil
}

func WriteZup(fileName string, zup Zup, key string) error {
	zupBytes, err := json.Marshal(zup)
	if err != nil {
		return err
	}

	encryptedContent, err := encrypt(string(zupBytes), []byte(key))
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, []byte(encryptedContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func AddFormatToZup(zup Zup, formatName string, fields []string) {
	zup.AddFormat(formatName, fields)
}

func GenerateZupKey() (string, error) {
	key, err := GenerateKey(KEY_SIZE)
	if err != nil {
		return "", err
	}

	readableKey := hex.EncodeToString(key)

	return readableKey, nil
}
