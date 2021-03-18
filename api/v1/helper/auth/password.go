package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	if res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err == nil {
		return string(res), nil
	} else {
		return "", err
	}
}

func Verify(password string, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err == nil {
		return true
	}
	return false
}
