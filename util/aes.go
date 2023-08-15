package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func Encrypt(text string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(ciphertext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertextBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertextBytes))
	stream.XORKeyStream(plaintext, ciphertextBytes)

	return string(plaintext), nil
}

func main() {
	key := []byte("example-key-1234")

	originalText := "Hello, encryption!"
	fmt.Println("Original Text:", originalText)

	encryptedText, err := Encrypt(originalText, key)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("Encrypted Text:", encryptedText)

	decryptedText, err := Decrypt(encryptedText, key)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("Decrypted Text:", decryptedText)
}
