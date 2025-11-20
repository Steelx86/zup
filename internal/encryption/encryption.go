package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

var (
	ErrInvalidKeySize = errors.New("invalid key size inputted")
	ErrShortCipher    = errors.New("ciphertext is too short")
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

func encryptZup(content string, key []byte) (string, error) {
	zupData := []byte(content)

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

func decryptZup(encryptedText string, key []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", fmt.Errorf("failed to decode: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", ErrShortCipher
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed ot decrypt data: %v", err)
	}

	content := string(plaintext)
	return content, nil
}
