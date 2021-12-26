package gopass

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(plain, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}
