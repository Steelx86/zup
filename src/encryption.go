package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidKeySize      = errors.New("Invalid key size inputted")
	ErrShortCipher         = errors.New("Ciphertext is too short")
)

func generateKey(size int) ([]byte, error) {
	if size != 16 && size != 24 && size != 32 {
		return nil, ErrInvalidKeySize
	}

	key := make([]byte, size)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func encryptZup(zupFile Zup, key []byte) (string, error) {
	zupData := []byte(zupFile.String())

	block, err := aes.NewCipher(zupData)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM %v", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce %v", err)
	}

	ciphertext := aesGCM.Seal(nil, nonce, zupData, nil)
	ciphertext = append(nonce, ciphertext...)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decryptZup(zupCipher string, key []byte) (Zup, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(zupCipher)
	if err != nil {
		return Zup{}, fmt.Errorf("failed to decode: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return Zup{}, fmt.Errorf("failed to create cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return Zup{}, fmt.Errorf("failed to create GCM: %v", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return Zup{}, ErrShortCipher
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return Zup{}, fmt.Errorf("failed ot decrypt data: %v", err)
	}

	zupFile := readZupString(string(plaintext))
	return zupFile, nil
}
