package model

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

type PassEntity struct {
	Name          string
	PassRawText   string
	PassEncrypted string
}

func New(entityName string, password string, encryptKey []byte) (PassEntity, error) {
	passEncrypted, err := Encrypt(password, encryptKey)
	if err != nil {
		return PassEntity{}, err
	}
	return PassEntity{
		Name:          entityName,
		PassRawText:   password,
		PassEncrypted: passEncrypted,
	}, nil
}

func (entry *PassEntity) GetPassEntryString() (string, error) {
	if entry.Name != "" && entry.PassEncrypted != "" {
		return entry.Name + " " + entry.PassEncrypted, nil
	} else {
		return "", errors.New("Invalid entry or password")
	}
}

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
