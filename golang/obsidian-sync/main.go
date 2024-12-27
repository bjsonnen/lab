package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return ciphertext, nil
}

func decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)
	return data, nil
}

func main() {
	password := "Password123!"
	message := []byte("Hello")

	hash := sha256.Sum256([]byte(password))

	data, err := encrypt(message, hash[:])
	if err != nil {
		fmt.Printf("Error encrypting data: %v\n", err)
		return
	}

	decryptedData, _ := decrypt(data, hash[:])

	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Hashed Password: %x\n", hash)
	fmt.Printf("Original Message: %s\n", message)
	fmt.Printf("Encrypted Data: %x\n", data)
	fmt.Printf("Decrypted Data: %s\n", decryptedData)
}
