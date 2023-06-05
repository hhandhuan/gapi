package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func CompareHashAndPassword(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
