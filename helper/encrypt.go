package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

var key = []byte("gwtq-sncb-siab-q")

func EncryptData(text string) string {
	plainText := []byte(text)
	block, err := aes.NewCipher(key)
	PanicIfError(err)

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		PanicIfError(err)
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.URLEncoding.EncodeToString(cipherText)
}

func DecryptData(encryptedStr string) string {
	cipherText, err := base64.URLEncoding.DecodeString(encryptedStr)
	PanicIfError(err)

	block, err := aes.NewCipher(key)
	PanicIfError(err)

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}
