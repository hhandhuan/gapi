package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

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
