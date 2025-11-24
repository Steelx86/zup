package zup

import (
	"encoding/json"
	"encoding/hex"
	"path/filepath"
	"fmt"
	"os"
	"strings"
)

const (
	KEY_SIZE = 32
	ZUP_DIR = "./.zup"
)

func InitZup(name string) (string, error) {
	if err := os.MkdirAll(ZUP_DIR, os.ModePerm); err != nil {
		return "", err
	}

	name = filepath.Clean(name)
	if strings.Contains(name, "..") || filepath.IsAbs(name) {
		return "", fmt.Errorf("Invalid file: %s")
	}

	if !strings.HasSuffix(name, ".zup") {
		name += ".zup"
	}

	filePath := filepath.Join(ZUP_DIR, name)
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create zup file: %v", err)
	}
	defer file.Close()

	key, err := GenerateKey(KEY_SIZE)
	if err != nil {
		return "", err
	}

	readableKey := hex.EncodeToString(key)
	return readableKey, nil
}

func OpenZup(fileName string, key string) (*Zup, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return &Zup{}, err
	}

	encryptedContent := string(file)

	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return &Zup{}, err
	}

	content, err := Decrypt(encryptedContent, &keyBytes)
	if err != nil {
		return &Zup{}, err
	}

	zupData, err := ParseZupString(content)
	if err != nil {
		return &Zup{}, err
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

	keyBytes, err := hex.DecodeString(key)
	if err != nil {
		return err
	}

	encryptedContent, err := Encrypt(string(zupBytes), &keyBytes)
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
