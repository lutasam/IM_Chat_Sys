package utils

import (
	"github.com/lutasam/chat/biz/common"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(s string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

func ValidatePassword(secret, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(secret), []byte(password))
	if err != nil {
		return common.PASSWORDISERROR
	}
	return nil
}
