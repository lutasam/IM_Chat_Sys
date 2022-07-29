package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/lutasam/chat/biz/common"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func Encrypt(s string) (string, error) {
	c, err := aes.NewCipher([]byte(common.PASSWORDSALT))
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(s))
	cfb.XORKeyStream(ciphertext, []byte(s))
	return string(ciphertext), nil
}

func Decrypt(secret string) (string, error) {
	c, err := aes.NewCipher([]byte(common.PASSWORDSALT))
	if err != nil {
		return "", err
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintext := make([]byte, len(secret))
	cfbdec.XORKeyStream(plaintext, []byte(secret))
	return string(plaintext), nil
}
