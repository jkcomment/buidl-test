package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
)

func Encrypt(key []byte, msg string) (string, error) {
	plainText := []byte(msg)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("##### failed create block: ", err)
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize + len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		log.Println("##### failed io.ReadFull: ", err)
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	encmess := base64.URLEncoding.EncodeToString(cipherText)

	return encmess, nil
}

func Decrypt(key []byte, msg string) (string, error) {
	cipherText, err := base64.URLEncoding.DecodeString(msg)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("##### Ciphertext size is too short!")
		return "", err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	decodedmess := string(cipherText)

	return decodedmess, nil
}
