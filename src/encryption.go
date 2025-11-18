package src

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func genKey(size uint8) ([]byte, error) {
	key := make([]byte, size)

	_, err := rand.Read(key); if err != nil {
		return nil, err
	}

	return key, err
}

func encryptZup(zupFile Zup, key []byte) (string, error) {
	zupData := []byte(string(zupFile.String()))
	block, err := aes.NewCipher(zupData); if err != nil {
		return "", err
	}

	// Random IV
	ciphertext := make([]byte, aes.BlockSize+len(zupData))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Data encryption
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.XORKeyStream(ciphertext[aes.BlockSize:], zupData)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decryptZup(zupCipher string, key []byte) (Zup, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(zupCipher); if err != nil {
		return Zup, err
	}


}
